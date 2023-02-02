from tba.types import Thread
import requests
import time
import xmltodict


def BackendLinkFlow(_, parentThread: Thread):
    print("[BACKENDLINK] Starting thread.")

    headers = {
        "authority": "www.the-broken-arm.com",
        "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8",
        "accept-language": "en-GB,en;q=0.7",
        "cache-control": "max-age=0",
        "cookie": "PHPSESSID=s2vjaar1s1hm5nvim1so9ro949; PrestaShop-b255acdcaf89d3f7cc8c1687088165cb=def5020075b6c74e228f544648ec9a547d7519a420b574216105675770664ce3123b5c3086a35cd89e24af1a06f2aba2e71a6c67cc85f257322b6d0809d6e4d3289d6f219ffe692ca852e8c17ad9e8fb15019894e4ba19f2d65dae4060c1faeda28317b1d178817b8ad964fc3625b0c293990ef70ef5b700c3c00fcd6cf428bbe69a4b5f703cb84585a2d4554448c39b3fb01c432bf4935ad4b0cc83d01426d7e8837d9c0ce62ae534b55037453aee89cb0c30507f6dc38bc1f0cb0e827f021a84867666950ec2646259c31f2dce3f2a1cccd0bb544ab17e7fe07327a1ffdb8da8334d1c419295384d411aa53ea97f7a21ab95d1902d2a980dd80f9f966f63fabdbb0fa6f8ddf01ee917609d4873ffa3c622578d4a3dbff265da2f48a0abae5e39053cf692687c75ba9cc01ed3; tarteaucitron=!analytics=true!gtag=true; cf_clearance=3TiED9b5DcA1rfD6mMJf.5ktsvE9ChwhGcLzKmWJ2Hw-1673965737-0-150; __cf_bm=QH.HhHUebksFqb_nEB50Ukg8uBS9LkCdDE.xhOGK1H8-1674421413-0-AXwugcPlxPn6MBBgW/alFH+1VryrsmPy0YDrlV++0SR6VbgX+oo0QfbCh76Ya5dP0EYPRzh+HqL3bx1bFNgDFs8tXYJcxxxvGz2Jnz1sLXEkIkHaoAh5rOPtfJz5pFQpLD7G8Pww5r16pbOiWH7b+vg=",
        "sec-ch-ua": '"Not_A Brand";v="99", "Brave";v="109", "Chromium";v="109"',
        "sec-ch-ua-mobile": "?0",
        "sec-ch-ua-platform": '"macOS"',
        "sec-fetch-dest": "document",
        "sec-fetch-mode": "navigate",
        "sec-fetch-site": "none",
        "sec-fetch-user": "?1",
        "sec-gpc": "1",
        "upgrade-insecure-requests": "1",
        "user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",
    }

    currentProducts = []

    while not parentThread.stop:
        try:
            time.sleep(1)
            try:
                response = requests.get(
                    "https://www.the-broken-arm.com/1_en_0_sitemap.xml",
                    headers=headers,
                )
            except requests.exceptions.ConnectionError or requests.exceptions.ConnectTimeout:
                print("[KEYWORD] Connection error.")
                continue

            if response.status_code == 200:
                try:
                    data = xmltodict.parse(response.text)
                except xmltodict.expat.ExpatError:
                    print("[BACKENDLINK] XML Error.")
                    continue

                print("[BACKENDLINK] Successfully fetched data.")

            else:
                print("[BACKENDLINK] Error fetching data.", response.status_code)

        except Exception as e:
            print("[KEYWORD] Error: {}".format(e))
            continue
