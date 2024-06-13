set
Charles > Proxy > Web Interface
1. Enable web interface
2. Allow anonymous access

get ip address by below
- ipconfig getifaddr en0
192.168.0.92
curl -v -x http://192.168.0.92:8888 "http://control.charles/throttling/activate?preset=3G"
curl -v -x http://192.168.0.92:8888 "http://control.charles/throttling/deactivate"


1. use default presets
2. set timer
3. write test scenario
4. add/remove presets
5. add/remove test scenario