# from dotenv import load_dotenv
import threading
from tba.start import Start
import sys

# load_dotenv()


def main():
    threading.Thread(target=Start).start()

    while True:
        q = input()
        if q.lower() in ["q", "quit"]:
            sys.exit(0)


if __name__ == "__main__":
    main()
