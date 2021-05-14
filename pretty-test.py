import json
import subprocess


def get_test_output():
    process = subprocess.Popen("go test -json -count=1 ./...", shell=True, stdin=subprocess.DEVNULL,
                               stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    return json.loads(b"[" + b", ".join(
        filter(lambda entry: len(entry) > 0, (process.stdout.read() + process.stderr.read()).split(b"\n"))) + b"]")


def compile_results(tests):
    packages = []
    for test in tests:
        if test["Action"] == "output":
            if test["Package"] not in packages:
                packages.append(test["Package"])
                r = " ".join(("Testing:", test["Package"] + ":", test["Output"]))
                print("-" * len(r))
                print(r, end="")
            else:
                if test.get("Test") is not None and test.get("Output") is not None:
                    print("\t\t", test["Output"], end="")
        elif test["Action"] == "run":
            print("\t", test["Test"])


def main():
    compile_results(get_test_output())


if __name__ == '__main__':
    main()
