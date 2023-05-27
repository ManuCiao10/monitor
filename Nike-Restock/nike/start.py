import requests
import time
from nike.types import headers
from nike.webhook import webhook_uzumaki, webhook_titan


def Start():
    print("[BACKEND] Starting thread...")

    firstRun = True
    shoesInfo: list[int] = {}
    array_size = ""

    sku_test = "FJ0704-100,DO6485-600"

    url = f"https://api.nike.com/product_feed/threads/v2?filter=language(it)&filter=marketplace(IT)&filter=channelId(d9a5bc42-4b9c-4976-858a-f159cf99c647)&filter=productInfo.merchProduct.styleColor({sku_test})"

    while True:
        try:
            print("[BACKEND] Request for products...")

            try:
                r = requests.get(url, headers=headers)
                data = r.json()

                print("[BACKEND] Successful got data.")
            except requests.exceptions.ConnectionError or requests.exceptions.ConnectTimeout:
                print("[BACKEND] Connection error.")
                continue

            except Exception as e:
                print("[BACKEND] Error: {}".format(e))
                continue

            if firstRun:
                firstRun = False

                for item in data["objects"]:
                    status = item["productInfo"][0]["merchProduct"]["status"]
                    styleColor = item["productInfo"][0]["merchProduct"]["styleColor"]
                    channels = item["productInfo"][0]["merchProduct"]["channels"]
                    imageUrls = item["productInfo"][0]["imageUrls"]["productImageUrl"]
                    fullPrice = item["productInfo"][0]["merchPrice"]["fullPrice"]
                    available = item["productInfo"][0]["availability"]["available"]
                    title = item["productInfo"][0]["productContent"]["title"]
                    slug = item["productInfo"][0]["productContent"]["slug"]

                    for item in item["productInfo"]:
                        try:
                            for size in item["skus"]:
                                id = size["id"]
                                localizedSize = size["countrySpecifications"][0][
                                    "localizedSize"
                                ]
                                for stock in item["availableSkus"]:
                                    if stock["id"] == id:
                                        availableSkus = stock["level"]
                                        array_size += (
                                            f"{localizedSize} [{availableSkus}]\n"
                                        )
                        except:
                            print("[BACKEND] Error: {}".format(styleColor))
                            array_size = ""

                        shoesInfo[styleColor] = {
                            "status": status,
                            "styleColor": styleColor,
                            "channels": channels,
                            "imageUrls": imageUrls,
                            "fullPrice": fullPrice,
                            "available": available,
                            "title": title,
                            "slug": slug,
                            "size": array_size,
                        }

                        array_size = ""

                print("[BACKEND] Saved data (first run).")
            else:
                print("[BACKEND] Checking for new products...")

                for item in data["objects"]:
                    styleColor = item["productInfo"][0]["merchProduct"]["styleColor"]
                    for item in item["productInfo"]:
                        try:
                            for size in item["skus"]:
                                id = size["id"]
                                localizedSize = size["countrySpecifications"][0][
                                    "localizedSize"
                                ]
                                for stock in item["availableSkus"]:
                                    if stock["id"] == id:
                                        availableSkus = stock["level"]
                                        array_size += (
                                            f"{localizedSize} [{availableSkus}]\n"
                                        )
                        except:
                            print(f"[BACKEND] [{styleColor}] error gettings sizes.")

                        # shoesInfo[styleColor]["size"] = "test"  # porpuse test only
                        if shoesInfo[styleColor]["size"] != array_size:
                            shoesInfo[styleColor]["size"] = array_size

                            print("[BACKEND] Size changed.")

                            webhook_uzumaki(shoesInfo[styleColor])
                            webhook_titan(shoesInfo[styleColor])

                        else:
                            print(f"[BACKEND] [{styleColor}] Size not changed.")
                        array_size = ""

                print("[BACKEND] Sleep for 1 seconds.")
                time.sleep(1)

        except Exception as e:
            print("[BACKEND] Error: {}".format(e))
            continue
