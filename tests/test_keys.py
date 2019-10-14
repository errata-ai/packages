import requests

from . import get_library


def test_keys():
    """Ensure that all required keys are defined.
    """
    for entry in get_library():
        assert list(entry.keys()) == [
            "name",
            "description",
            "homepage",
            "url"
        ]


def test_links():
    """Ensure that our all links are working.
    """
    for entry in get_library():
        for k in ("homepage", "url"):
            link = entry[k]
            r = requests.head(link)
            assert r.status_code in (200, 302), link
