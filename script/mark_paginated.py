#!/usr/bin/env python3

# This script isn't run from Makefile. It's here to help generate patch for the
# x-paginated object. The x-paginated object marks that all-pages-listing 
# method should be generated for given operation, think listing of devices
# in a project with more than 20 devices.

# example run:
# for o in `find spec/oas3.patched/paths -name '*.yaml'`; do script/mark_paginated.py $o; done

import sys
import os
import os.path
import ruamel.yaml

yaml = ruamel.yaml.YAML()
yaml.preserve_quotes = True

ymlfile = sys.argv[1]

def getYaml(fn):
    with open(fn) as f:
        return yaml.load(f.read())

paginated = False
paginatedProperty = None

operation = getYaml(ymlfile)

def toGolangPropertyName(s):
    return "".join([w.capitalize() for w in s.split("_")])

refkey = "$ref"
if "get" in operation:
    getop = operation['get']
    if "parameters" in getop:
        getopparams = getop['parameters']
        for p in getopparams:
            if refkey in p:
                if ("/components/parameters/Page.yaml" in p[refkey]) :
                    schemaPathRelative = getop['responses']['200']['content']['application/json']['schema']['$ref']
                    schemaPath = schemaPathRelative.split("../")[-1]
                    twoTopDirs = ymlfile.split("/")[:2]
                    specDir = os.path.dirname(ymlfile)
                    schemaFile = os.path.join(specDir, schemaPathRelative)
                    responseSchema = getYaml(schemaFile)
                    responseProps = set(responseSchema['properties'].keys())
                    oid = getop['operationId']
                    if 'meta' in responseProps:
                        responseProps.remove("meta")
                    else:
                        print("%s => %s doesn't have 'meta', and therefore no PageNum etc" % (oid, schemaFile))
                        break
                    if len(responseProps) > 1:
                        print("%s => %s has 'page' but ambiguous response type with args %s" %
                              (oid, schemaFile, responseProps))
                        break
                    print("marking %s as paginated" % oid)
                    paginated = True
                    paginatedProperty = toGolangPropertyName(responseProps.pop())
                    break


if paginated:
    operation['get']['x-equinix-metal-paginated-property'] = paginatedProperty
    with open(ymlfile,"w") as f:
        yaml.dump(operation, f)


