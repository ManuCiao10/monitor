import tls_client
import time
from nike.webhook import webhook


headers = {
    "authority": "www.nike.com.ar",
    "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8",
    "accept-language": "en-GB,en-US;q=0.9,en;q=0.8",
    "cache-control": "no-cache",
    "pragma": "no-cache",
    "sec-ch-ua": '"Not.A/Brand";v="8", "Chromium";v="114", "Brave";v="114"',
    "sec-ch-ua-mobile": "?0",
    "sec-ch-ua-platform": '"macOS"',
    "sec-fetch-dest": "document",
    "sec-fetch-mode": "navigate",
    "sec-fetch-site": "none",
    "sec-fetch-user": "?1",
    "sec-gpc": "1",
    "service-worker-navigation-preload": "true",
    "upgrade-insecure-requests": "1",
    "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
}

params = {
    "workspace": "master",
    "maxAge": "short",
    "appsEtag": "remove",
    "domain": "store",
    "locale": "es-AR",
    "__bindingId": "d69fd38e-69f0-40e3-a6cb-27082d8c65ff",
    "operationName": "productSearchV3",
    "variables": "{}",
    "extensions": '{"persistedQuery":{"version":1,"sha256Hash":"40e207fe75d9dce4dfb3154442da4615f2b097b53887a0ae5449eb92d42e84db","sender":"vtex.store-resources@0.x","provider":"vtex.search-graphql@0.x"},"variables":"eyJoaWRlVW5hdmFpbGFibGVJdGVtcyI6dHJ1ZSwic2t1c0ZpbHRlciI6IkZJUlNUX0FWQUlMQUJMRSIsInNpbXVsYXRpb25CZWhhdmlvciI6ImRlZmF1bHQiLCJpbnN0YWxsbWVudENyaXRlcmlhIjoiTUFYX1dJVEhPVVRfSU5URVJFU1QiLCJwcm9kdWN0T3JpZ2luVnRleCI6ZmFsc2UsIm1hcCI6InByb2R1Y3RDbHVzdGVySWRzLHRpcG8tZGUtcHJvZHVjdG8scHJvZHVjdGNsdXN0ZXJuYW1lcyIsInF1ZXJ5IjoiMTcyL2NhbHphZG8vbGFuemFtaWVudG9zIiwib3JkZXJCeSI6Ik9yZGVyQnlSZWxlYXNlRGF0ZURFU0MiLCJmcm9tIjoxOCwidG8iOjM1LCJzZWxlY3RlZEZhY2V0cyI6W3sia2V5IjoicHJvZHVjdENsdXN0ZXJJZHMiLCJ2YWx1ZSI6IjE3MiJ9LHsia2V5IjoidGlwby1kZS1wcm9kdWN0byIsInZhbHVlIjoiY2FsemFkbyJ9LHsia2V5IjoicHJvZHVjdGNsdXN0ZXJuYW1lcyIsInZhbHVlIjoibGFuemFtaWVudG9zIn1dLCJzZWFyY2hTdGF0ZSI6bnVsbCwiZmFjZXRzQmVoYXZpb3IiOiJTdGF0aWMiLCJjYXRlZ29yeVRyZWVCZWhhdmlvciI6ImRlZmF1bHQiLCJ3aXRoRmFjZXRzIjpmYWxzZX0="}',
}


def Start():
    print("[BACKEND] Starting thread...")

    firstRun = True
    firstID = None
    while True:
        try:
            try:
                session = tls_client.Session(
                    client_identifier="chrome114", random_tls_extension_order=True
                )
                print("[BACKEND] Getting data...")
                response = session.get(
                    "https://www.nike.com.ar/_v/segment/graphql/v1",
                    params=params,
                    headers=headers,
                )
                data = response.json()
                print("[BACKEND] Successful got data.")

            except Exception as e:
                print("[BACKEND] Error getting data: {}".format(e))
                time.sleep(5)
                continue

            if firstRun:
                firstID = data["data"]["productSearch"]["products"][0]["productId"]
                print("[BACKEND] Got First-ID (first run): {}".format(firstID))
                firstRun = False
            else:
                print("[BACKEND] Checking for new IDs...")

                id = data["data"]["productSearch"]["products"][0]["productId"]
                # id = 1496 #testing purposes
                if firstID != id:
                    firstID = id
                    print("[BACKEND] Got New-ID: {}".format(id))
                    productName = data["data"]["productSearch"]["products"][0]["productName"]
                    linkText = data["data"]["productSearch"]["products"][0]["link"]
                    sku = data["data"]["productSearch"]["products"][0]["productReference"]
                    image = data["data"]["productSearch"]["products"][0]["items"][0]["images"][0]["imageUrl"]
                    availableQuantity = data["data"]["productSearch"]["products"][0]["items"][0]["sellers"][0]["commertialOffer"]["AvailableQuantity"]
                    listPrice = data["data"]["productSearch"]["products"][0]["items"][0]["sellers"][0]["commertialOffer"]["ListPrice"]

                    for i in range(len(data["data"]["productSearch"]["products"][0]["specificationGroups"][1]["specifications"])):
                        if data["data"]["productSearch"]["products"][0]["specificationGroups"][1]["specifications"][i]["name"] == "talle":
                            talle = data["data"]["productSearch"]["products"][0]["specificationGroups"][1]["specifications"][i]["values"]
                            break

                    webhook(productName,linkText,sku,image,availableQuantity,listPrice,talle)
                else:
                    print("[BACKEND] No new IDs found.")

                print("[BACKEND] Sleeping for 300 seconds...")
                time.sleep(60)

        except Exception as e:
            print("[BACKEND] Error: {}".format(e))
            time.sleep(60)
            continue



