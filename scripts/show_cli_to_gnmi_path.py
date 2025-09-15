import shlex
import sys
import argparse

NO_COMMAND_ERROR = "No command provided for conversion"
EMPTY_COMMAND_ERROR = "Empty command"
INVALID_COMMAND_ERROR = "Command must start with 'show'"
NO_PATH_ERROR = "No path segments after 'show'"
SHORT_OPTION_ERROR = "Short options are not supported"
INVALID_OPTION_TOKEN_ERROR = "Invalid option token: --"
INVALID_LONG_OPTION_ERROR = "Invalid long option '--'"

ESCAPABLE = frozenset({"/", "[", "]", "=", "\\"})


class OptionException(Exception):
    """Raised for option/command validation errors (not parsing syntax)."""
    pass


def escape_gnmi(text: str) -> str:
    """
    Escape gNMI special characters using a single ESCAPABLE set:
      - If '\' precedes an ESCAPABLE char, keep that escape as-is.
      - Else a lone '\' becomes '\\\\'.
      - Any raw / [ ] = get a leading backslash.
    """
    out = []
    i, n = 0, len(text)

    while i < n:
        ch = text[i]

        if ch == "\\":
            if i + 1 < n and text[i + 1] in ESCAPABLE:
                # escaping / [ ] = \
                out.append("\\")
                out.append(text[i + 1])
                i += 2
            else:
                # Literal backslash
                out.append("\\\\")
                i += 1
            continue

        if ch in ESCAPABLE:
            out.append("\\" + ch)
        else:
            out.append(ch)

        i += 1

    return "".join(out)


class ShowCliToGnmiPathConverter:
    def __init__(self, command: str):
        self.command = command

    def parseLongOption(self, token: str):
        # --flag         -> ('flag', 'True')
        # --key=value    -> ('key',  'value')
        if token == "--":
            raise OptionException(INVALID_OPTION_TOKEN_ERROR)

        body = token[2:]
        if not body:
            raise OptionException(INVALID_LONG_OPTION_ERROR)

        if "=" in body:
            name, value = body.split("=", 1)
            if not name:
                raise OptionException("Invalid long option: missing name before '='")
            return name, escape_gnmi(value)

        return body, "True"

    def convert(self) -> str:
        tokens = shlex.split(self.command)

        if not tokens:
            raise OptionException(EMPTY_COMMAND_ERROR)
        if tokens[0].lower() != "show":
            raise OptionException(INVALID_COMMAND_ERROR)

        tokens = tokens[1:]  # drop 'show'
        out = []

        for tok in tokens:
            # Reject short options
            if tok.startswith("-") and not tok.startswith("--"):
                raise OptionException(f"{SHORT_OPTION_ERROR}: '{tok}'")

            if tok.startswith("--"):
                # Option before any paths
                if not out:
                    raise ValueError("Option before first path segment")
                key, val = self.parseLongOption(tok)
                out.append(f"[{key}={val}]")
                continue

            if out:
                out.append("/")
            out.append(escape_gnmi(tok))

        if not out:
            # No segments at all
            raise OptionException(NO_PATH_ERROR)

        return "".join(out)


def main():
    parser = argparse.ArgumentParser(add_help=True)
    parser.add_argument(
        "-c",
        "--command",
        nargs="?",
        const="",
        default="",
        help='Command string starting with "show", e.g. "show foo bar --key=value"',
    )
    args = parser.parse_args()

    if args.command == "":
        print(NO_COMMAND_ERROR, file=sys.stderr)
        sys.exit(2)

    try:
        result = ShowCliToGnmiPathConverter(args.command).convert()
    except OptionException as e:
        print(str(e), file=sys.stderr)
        sys.exit(1)
    except ValueError as e:
        print(f"failed to parse command: {e}", file=sys.stderr)
        sys.exit(1)

    print(result)


if __name__ == "__main__":
    main()
