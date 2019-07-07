from . import get_library


def test_json():
    """Ensure that our library is valid JSON.
    """
    get_library()


def test_order():
    """Ensure that our library is sorted alphabetically.
    """
    names = [e["name"] for e in get_library()]
    assert names == sorted(names)
