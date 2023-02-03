import threading
from airness.start import Start
from airness.util import PURPLE, RESET, BANNER


def main():
    """Main function"""
    print(PURPLE + BANNER + RESET)

    print("[*] Scraping airness.eu website...")
    threading.Thread(target=Start).start()


if __name__ == "__main__":
    main()
