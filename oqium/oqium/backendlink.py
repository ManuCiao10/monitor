import requests
import time
from oqium.typesOqium import Thread


def webhook(price, title, url_name, url_image, array_info, sku):
    url = ""

    embee = []
    embee.append({"name": "SKU", "value": sku, "inline": False})
    embee.append({"name": "Price", "value": "€" + str(price), "inline": False})

    for i in array_info:
        embee.append(
            {"name": "Size", "value": i[0] + " - [" + str(i[1]) + "]", "inline": True}
        )

    data = {
        "username": "Oqium Monitor",
        "avatar_url": "",
        "embeds": [
            {
                "title": title,
                "url": "https://oqium.com/collections/footwear-1/products/" + url_name,
                "color": 15794176,
                "thumbnail": {"url": url_image},
                "fields": embee,
                "footer": {
                    "text": "Oqium Monitor",
                    "icon_url": "",
                },
            }
        ],
    }

    result = requests.post(url, json=data)
    try:
        result.raise_for_status()
    except requests.exceptions.HTTPError as err:
        print(err)


def BackendLinkFlow(_, parentThread: Thread):
    print("[BACKENDLINK] Starting thread.")

    headers = {
        "Accept": "*/*",
        "Accept-Language": "en-GB,en-US;q=0.9,en;q=0.8",
        "Cache-Control": "no-cache",
        "Connection": "keep-alive",
        "Content-Type": "application/x-www-form-urlencoded",
        "Origin": "https://oqium.com",
        "Pragma": "no-cache",
        "Referer": "https://oqium.com/",
        "Sec-Fetch-Dest": "empty",
        "Sec-Fetch-Mode": "cors",
        "Sec-Fetch-Site": "cross-site",
        "Sec-GPC": "1",
        "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
        "sec-ch-ua": '"Chromium";v="110", "Not A(Brand";v="24", "Brave";v="110"',
        "sec-ch-ua-mobile": "?0",
        "sec-ch-ua-platform": '"macOS"',
    }

    params = {
        "q": "",
        "apiKey": "d5498f0b-2aa4-4b7d-b7d1-ffd7d8761ec1",
        "locale": "en",
        "getProductDescription": "0",
        "collection": "162173321312",
        "skip": "0",
        "take": "1",
        "sort": "-date",
    }

    sku_array = []
    first_run = False

    while True:
        try:
            time.sleep(5)
            response = requests.get(
                "https://svc-0-usf.hotyon.com/search", params=params, headers=headers
            )

            try:
                response_data = response.json()
            except ValueError:
                print("[BACKENDLINK] Error: Invalid JSON")
                time.sleep(5)
                continue

            try:
                product_data = response_data["data"]["items"][0]
                sku = product_data["tags"][0]
                title = product_data["title"]
                url_name = product_data["urlName"]
                url_image = "https:" + product_data["images"][0]["url"]
                variants = product_data["variants"]
                options = product_data["options"][0]
                values_dict = {
                    i: value for i, value in enumerate(options.get("values", []))
                }
                array_info = [
                    [
                        values_dict[variant["options"][0]].split("|")[1].strip(),
                        variant["available"],
                    ]
                    for variant in variants
                ]
                price = variants[0]["price"]

                if sku not in sku_array:
                    sku_array.append(sku)
                    print("[BACKENDLINK] Adding " + sku + "...")

                    if first_run:
                        print("[BACKENDLINK] New product found...")
                        webhook(price, title, url_name, url_image, array_info, sku)

                    first_run = True

                else:
                    time_ = time.strftime("%H:%M:%S", time.localtime())
                    print(time_, "[BACKENDLINK] No new products found...")
            except KeyError:
                time.sleep(20)
                time_ = time.strftime("%H:%M:%S", time.localtime())
                print(time_, "[BACKENDLINK] Rate limited...")
                continue

        except requests.exceptions.RequestException as e:
            print("[BACKENDLINK] Error: " + str(e))
            time.sleep(5)
            return
