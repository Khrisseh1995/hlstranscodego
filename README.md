# Serverside HLS Ad Insertion using Golang
This service will (hopefully) achieve serverside ad insertion in HLS Vod Streams

The full panthos 23 spec can be found [here](https://tools.ietf.org/html/draft-pantos-http-live-streaming-23)


## List of example playlists

I'm currently using playlists from BitMovin hosted [here](https://bitmovin.com/mpeg-dash-hls-examples-sample-streams/) as I currently only have working implementations for video and audio, the best examples to use are:

If you're looking at HLS streams on GitHub you most likely know this. These files will only playback natively in Safari

- https://bitdash-a.akamaihd.net/content/MI201109210084_1/m3u8s/f08e80da-bf1d-4e3d-8899-f0f6155f6efa.m3u8

## Example Call

`http://localhost:8080/generate_master_playlist?masterPlaylist=https://bitdash-a.akamaihd.net/content/MI201109210084_1/m3u8s/f08e80da-bf1d-4e3d-8899-f0f6155f6efa.m3u8&baseUrl=https://bitdash-a.akamaihd.net/content/MI201109210084_1/m3u8s` 


