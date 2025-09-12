import re
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

_UNESCAPED_SLASH = re.compile(r"(?<!\\)/")


class CliToXpathConverter:
    def __init__(self, command):
        self.command = command

    def escapeSlashes(self, text):
        return _UNESCAPED_SLASH.sub(r"\/", text)

    def raiseError(self, message):
        raise Exception(message)

    def parseLongOption(self, token):
        # --flag         -> ('flag', 'True')
        # --key=value    -> ('key',  'value')
        if token == "--":
            self.raiseError(INVALID_OPTION_TOKEN_ERROR)

        body = token[2:]
        if not body:
            self.raiseError(INVALID_LONG_OPTION_ERROR)

        if "=" in body:
            name, value = body.split("=", 1)
            if not name:
                self.raiseError("Invalid long option: missing name before '='")
            return name, self.escapeSlashes(value)

        return body, "True"

    def convert(self):
        try:
            tokens = shlex.split(self.command)
        except Exception as e:
            self.raiseError(f"failed to parse command: {e}")

        if not tokens:
            self.raiseError(EMPTY_COMMAND_ERROR)

        if tokens[0].lower() != "show":
            self.raiseError(INVALID_COMMAND_ERROR)

        tokens = tokens[1:]  # drop 'show'

        path_segments = []
        option_blocks = []

        i = 0
        n = len(tokens)
        while i < n:
            tok = tokens[i]

            # Reject bare '-' or any short option form
            if tok == "-" or (tok.startswith("-") and not tok.startswith("--")):
                self.raiseError(f"{SHORT_OPTION_ERROR}: '{tok}'")

            if tok.startswith("--"):
                key, val = self.parseLongOption(tok)
                option_blocks.append(f"[{key}={val}]")
                i += 1
                continue

            # Positional => path segment
            path_segments.append(self.escapeSlashes(tok))
            i += 1

        if not path_segments:
            self.raiseError(NO_PATH_ERROR)

        return "/".join(path_segments) + "".join(option_blocks)


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
        result = CliToXpathConverter(args.command).convert()
    except Exception as e:
        print(str(e), file=sys.stderr)
        sys.exit(1)

    print(result)


if __name__ == "__main__":
    main()
