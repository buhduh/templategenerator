#!/usr/bin/env python3

import sys
import json

def run():
    data = {}
    with open(sys.argv[1]) as pkgJSON:
        data = json.load(pkgJSON)
        data["version"] = sys.argv[2]
    with open(sys.argv[1], "w") as pkgJSON:
        json.dump(data, pkgJSON, indent=2)

if __name__ == '__main__':
    run()
