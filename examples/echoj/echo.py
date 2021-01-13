#!/usr/bin/env python2.7

import json
import sys
import time


def cycleWork():
	oneline = sys.stdin.readline()
	obj_in = json.loads(oneline)
	obj_out = {
			"input": obj_in,
			"cmd": sys.argv,
			"clock": time.time(),
	}
	json.dump(obj_out, sys.stdout)
	sys.stdout.write("\n")
	sys.stdout.flush()
	return obj_in


def main():
	with open("/tmp/123.log", "w") as fp:
		while True:
			o = cycleWork()
			fp.write(repr(o) + "\n")
	return 0


if __name__ == '__main__':
	sys.exit(main())
