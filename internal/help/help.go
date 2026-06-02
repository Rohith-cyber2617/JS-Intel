package help

const HelpMenu = `
NAME:
    JS Intel

DESCRIPTION:
    JavaScript Intelligence Gathering Framework

USAGE:
    jsintel [OPTIONS]

OPTIONS:

    -u, --url URL
            Scan a single target URL.

    -l, --list FILE
            Scan multiple targets from a file.

    -o, --output FILE
            Save output to a file.

            Supported:
                .txt
                .json
                .xml
                .html

            Default:
                .txt

    -t, --threads INT
            Number of concurrent threads.

            Default: 10

    --depth INT
            Analysis depth.

            1 = JS Discovery
            2 = Endpoints + Parameters
            3 = Secrets + Tokens
            4 = Internal APIs + GraphQL
            5 = Full Intelligence Mode

            Default: 1

    -ep, --end-point
            Extract endpoints from
            JavaScript files.

    -fo, --found-only
            Display only interesting
            findings.

    --verify
            Verify findings.

    --random-agent
            Use random User-Agent.

    --update
    --up
            Update JS Intel.

    -h, --help
            Display this help menu.

EXAMPLES:

    jsintel -u https://target.com

    jsintel -u https://target.com --depth=5

    jsintel -l targets.txt

    jsintel -u https://target.com -o output.json

    jsintel -u https://target.com --verify
`
