## Summary

This script downloads videos from a closed course website. It requires the URLs of the videos, the cookies and headers of the requests, and a CSV file with the metadata of the videos. It downloads the videos.

I used this to download: https://bojtiandrea.hu/kistestver-szuletik-zart-kurzusfelulet/

This chat with ChatGPT helped me create this script: https://chatgpt.com/share/d557da42-4d3c-41d5-a6c2-03f81df2fcc7

## Usage
1. Go to page
2. Go to Dev Tools → Network → Fetch/xhr, filter for "master"
3. Hover through all videos in order to load them
4. Copy the URL of each request (quite tedious, I know), into the input/urls.txt. Mind that they'll expire in an hour.
5. Replace "master.json?..." with "master.mpd" in the URLs.
6. Select, then copy the complete page content and paste it to chatgpt, ask it to convert it to a CSV with 4 columns: 1. Index, 2. Filename (inferred from the title, but somewhat shorter), prefixed with the index, all lowercase, separated by underscore, and with the extension "mp4", 3. Title, without "fejezet", and no all caps, 4. Description (can be multiline)
7. Get the saved CSV and save it as input/chapters.csv, without any header.
8. Copy-paste your cookies into the input/cookies.txt
9. Copy-paste your headers for any of the "master" requests into the input/headers.txt
10. Run the script
11. Optionally, go to the Network tab again, and search for "vtt", then download the subtitles one by one. I haven't found a way to automate this, or even to find their correct order automatically, so it's a bit of a mess.