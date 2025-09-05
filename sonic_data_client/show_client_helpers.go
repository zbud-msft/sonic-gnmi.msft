package client

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	gnmipb "github.com/openconfig/gnmi/proto/gnmi"
	spb "github.com/sonic-net/sonic-gnmi/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func showHelp(prefix, path *gnmipb.Path, description map[string]map[string]string) ([]*spb.Value, error) {
	helpData, err := json.Marshal(description)
	if err != nil {
		return nil, err
	}

	var values []*spb.Value
	ts := time.Now()
	values = append(values, &spb.Value{
		Prefix:    prefix,
		Path:      path,
		Timestamp: ts.UnixNano(),
		Val: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_JsonIetfVal{
				JsonIetfVal: helpData,
			}},
	})
	return values, nil
}

func (spcfg ShowPathConfig) ParseOptions(path *gnmipb.Path) (OptionMap, error) {
	passedOptions, err := checkOptionsInPath(path, spcfg.options)
	if err != nil {
		return nil, err
	}
	return validateOptions(passedOptions, spcfg.options)
}

func (spcfg ShowPathConfig) ParseArgs(path *gnmipb.Path) (CmdArgs, error) {
	pathArr := pathToArr(path)
	if spcfg.maxArgs < -1 {
		return nil, status.Errorf(codes.Internal, "invalid number of max args: must be greater or equal to -1 (any # of args): %d", spcfg.maxArgs)
	}
	if spcfg.regLen <= 0 {
		return nil, status.Errorf(codes.Internal, "invalid config: registered prefix length: %d", spcfg.regLen)
	}
	argStartIndex := spcfg.regLen - 1 // args start after registered prefix
	if argStartIndex < 0 || argStartIndex > len(pathArr) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid path: expected atleast %d elements after target, got: %d", spcfg.regLen - 1, len(pathArr))
	}
	numArgs := len(pathArr) - argStartIndex
	if spcfg.maxArgs >= 0 && numArgs > spcfg.maxArgs {

		return nil, status.Errorf(codes.InvalidArgument, "invalid number of arguments provided: must be less than or equal to %d", spcfg.maxArgs)
	}
	return CmdArgs(pathArr[argStartIndex:]), nil
}

func pathToArr(path *gnmipb.Path) []string {
	out := make([]string, 0)
	if path != nil {
		elems := path.GetElem()
		for _, elem := range elems {
			out = append(out, elem.GetName())
		}
	}
	return out
}

func validateOptions(passedOptions map[string]string, options map[string]ShowCmdOption) (OptionMap, error) {
	optionMap := make(OptionMap)
	// Validate that mandatory options exist and unimplemented options are errored out and validate proper typing for each option
	for optionName, optionCfg := range options {
		optionValue, found := passedOptions[optionName]
		if !found {
			if optionCfg.optType == Required {
				return nil, status.Errorf(codes.InvalidArgument, "option %v is required", optionName)
			}
			continue
		}
		if optionCfg.optType == Unimplemented {
			return nil, status.Errorf(codes.Unimplemented, "option %v is unimplemented", optionName)
		}

		switch optionCfg.valueType {
		case StringValue:
			optionMap[optionName] = OptionValue{value: optionValue}
		case StringSliceValue:
			valueParts := strings.Split(optionValue, ",")
			for i := range valueParts {
				valueParts[i] = strings.TrimSpace(valueParts[i])
			}
			optionMap[optionName] = OptionValue{value: valueParts}
		case BoolValue:
			boolValue, err := strconv.ParseBool(optionValue)
			if err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "option %v expects a bool (got %v), err: %v", optionName, optionValue, err)
			}
			optionMap[optionName] = OptionValue{value: boolValue}
		case IntValue:
			intValue, err := strconv.Atoi(optionValue)
			if err != nil {
				return nil, status.Errorf(codes.InvalidArgument, "option %v expects an int (got %v), err: %v", optionName, optionValue, err)
			}
			optionMap[optionName] = OptionValue{value: intValue}
		case EnumValue:
			valid := false
			for _, v := range optionCfg.enumValues {
				if v == optionValue {
					valid = true
					break
				}
			}
			if !valid {
				return nil, status.Errorf(codes.InvalidArgument, "option %v expects one of [%v] (got %v)", optionName, strings.Join(optionCfg.enumValues, ", "), optionValue)
			}
			optionMap[optionName] = OptionValue{value: optionValue}
		default:
			return nil, status.Errorf(codes.InvalidArgument, "unsupported ValueType for option %v", optionName)
		}
	}
	return optionMap, nil
}

func checkOptionsInPath(path *gnmipb.Path, options map[string]ShowCmdOption) (map[string]string, error) {
	// Validate that path doesn't contain any option that is not registered
	passedOptions := make(map[string]string)
	for _, elem := range path.GetElem() {
		for key, val := range elem.GetKey() {
			if _, ok := options[key]; !ok {
				return nil, status.Errorf(codes.InvalidArgument, "option %v for path %v is not a valid option", key, path)
			}
			passedOptions[key] = val
		}
	}
	return passedOptions, nil
}

func constructDescription(usage string, subcommandDesc map[string]string, options map[string]ShowCmdOption) map[string]map[string]string {
	description := make(map[string]map[string]string)
	description["options"] = make(map[string]string)
	description["usage"] = make(map[string]string)

	for _, option := range options {
		// Base description
		desc := option.description

		// If option is EnumValue, append allowed values to the description
		if option.valueType == EnumValue && len(option.enumValues) > 0 {
			desc = fmt.Sprintf("%s (Allowed values: %s)", desc, strings.Join(option.enumValues, ", "))
		}

		description["options"][option.optName] = desc
	}

	description["subcommands"] = subcommandDesc
	description["usage"]["desc"] = usage
	return description
}

func constructOptions(options []ShowCmdOption) map[string]ShowCmdOption {
	pathOptions := make(map[string]ShowCmdOption)
	pathOptions[showCmdOptionHelp.optName] = showCmdOptionHelp
	for _, option := range options {
		pathOptions[option.optName] = option
	}
	return pathOptions
}
