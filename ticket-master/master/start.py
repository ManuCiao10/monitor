from master.masterTypes import Thread
from master.backendlink import BackendLinkFlow
from master.threads import RunningThreads


def Start():
    TICKET = [
        "https://shop.ticketmaster.it/biglietti/acquista-biglietti-travis-scott-i-days-milano-coca-cola-30-giugno-2023-ippodromo-snai-la-maura-milano-4996.html",
    ]

    while True:
        for pid in TICKET:
            if pid not in RunningThreads:
                RunningThreads[pid] = Thread(BackendLinkFlow, pid)
                RunningThreads[pid].start()
