import collections
import json


def get_library():
    """Read and return our current library.
    """
    with open("library.json") as f:
        return json.load(f, object_pairs_hook=collections.OrderedDict)
