#! /usr/local/bin/python3
"Installer script."


import argparse
import json
import logging
import os
import platform
import sys
import tarfile
import time
import warnings

import urllib3  # type: ignore

warnings.filterwarnings("ignore")

OWNER = "kislerdm"
APP = "github-downloader"
DIR = os.path.join(os.path.expanduser("~"), f".{APP}")
LINK = f"/usr/local/bin/{APP}"


class HTTP:
    def __init__(self):
        logging.getLogger("urllib3").setLevel(logging.WARNING)
        self.client = urllib3.PoolManager(cert_reqs="CERT_NONE")

    def get(self, url, headers=None):
        resp = self.client.request("GET", url, headers=headers)
        return resp.data

    def post(self, url, payload, headers=None):
        resp = self.client.request("POST", url, fields=payload, headers=headers)
        return resp.data


def json_unmarshal(obj):
    return json.loads(obj.decode("utf-8"))


def get_args():
    args = argparse.ArgumentParser(f"{APP} installation script")
    args.add_argument("-version", "-v", default="", required=False, help="Specific app version.")
    return args.parse_args()


def uname():
    system_info = platform.uname()
    return system_info.system.lower(), system_info.machine


def parse_ts_to_int(s):
    ts_local = time.strptime(s, "%Y-%m-%dT%H:%M:%SZ")
    return int(time.mktime(ts_local))


def get_app_releases(http_client):
    resp = json_unmarshal(
        http_client.get(
            url=os.path.join("https://api.github.com/repos", OWNER, APP, "releases"),
            headers={"Accept": "application/vnd.github.v3+json"},
        ),
    )
    return {parse_ts_to_int(r.get("published_at")): r.get("tag_name") for r in resp}


def get_last_release(releases):
    return releases[max(releases.keys())]


def get_download_url(version):
    os_ver, arch = uname()
    return os.path.join(
        "https://github.com",
        OWNER,
        APP,
        "releases/download",
        version,
        f"{APP}-{version}-{os_ver}-{arch}.tar.gz",
    )


def mkdir():
    if not os.path.isdir(DIR):
        os.makedirs(DIR)


def main(version):
    log = logging.getLogger(__name__)

    http_client = HTTP()
    app_releases = get_app_releases(http_client)

    if version != "":
        if version not in app_releases.values():
            raise ValueError(f"version {version} is not found in the repo. Please check the input.")
    else:
        version = get_last_release(app_releases)

    log.info(f"Downloading the app, version: {version}")

    app_release_url = get_download_url(version)
    app_archive_obj = http_client.get(app_release_url)

    log.info("Saving archive")

    mkdir()
    path_arch = f"{DIR}/{APP}-{version}.tar.gz"
    with open(path_arch, "wb") as fout:
        fout.write(app_archive_obj)

    log.info("Extracting files from archive")

    with tarfile.open(path_arch) as tar:
        tar.extractall(DIR)

    log.info(f"Linking the binary {os.path.dirname(LINK)}")

    if os.path.islink(LINK):
        log.info(f"The binary is linked to {LINK} already")
        answer = input("Would you like to overwrite? [Y/n] ")
        if answer in ("Y", "y"):
            os.system(f"sudo rm -r {LINK}")
        else:
            log.info("Installation interrupted")
            return

    os.system(f"sudo ln -sf {DIR}/{APP} {LINK}")
    return


if __name__ == "__main__":
    logging.basicConfig(
        level=logging.INFO,
        format=("[%(levelname)s] %(asctime)s %(message)s"),
        datefmt="%Y/%m/%d %H:%M:%S",
    )
    log = logging.getLogger(__name__)
    args = get_args()
    try:
        main(version=args.version)
    except Exception as ex:
        log.error(ex)
        sys.exit(1)
    log.info("Done. Bye!")
