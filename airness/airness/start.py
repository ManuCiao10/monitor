from airness.types import Thread
from airness.backend import BackendLinkFlow


def Start():
    """
    Scrape products from airness.eu website by
    using the BackendLinkFlow function
    """

    t = Thread(BackendLinkFlow, None)
    t.start()
