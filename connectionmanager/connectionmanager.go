package connectionmanager

import (
	"net/http"
	"encoding/json"
	"github.com/VladislavBryukhanov/voip-signaling/model"
)

func GetActiveConnections(w http.ResponseWriter, r *http.Request) {
	res := model.GetActiveConnections()

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func UpsertConnection(w http.ResponseWriter, r *http.Request) {
	mockPeerCon := &model.WebRTCConnection {
		InitiatorId: 1619334574997,
		ExpirationDate: 1619334623067,
		Offer: model.RTCSessionDescriptionInit {
			Sdp: "v=0 o=- 6615604271789689250 2 IN IP4 127.0.0.1 s=- t=0 0 a=group:BUNDLE 0 1 a=extmap-allow-mixed a=msid-semantic: WMS 9Frw0xYyaszH0JHrspJgRLPEs6eczFMDpBtt m=audio 9 UDP/TLS/RTP/SAVPF 111 103 104 9 0 8 106 105 13 110 112 113 126 c=IN IP4 0.0.0.0 a=rtcp:9 IN IP4 0.0.0.0 a=ice-ufrag:zh8p a=ice-pwd:95ldnP+KE1KcFetSAHJaqj5Z a=ice-options:trickle a=fingerprint:sha-256 82:DB:EA:2B:C4:BA:BD:30:55:E5:A4:56:21:EE:B5:92:93:D6:CB:BA:DB:E5:DD:05:B9:F3:97:BC:AE:48:2C:98 a=setup:actpass a=mid:0 a=extmap:1 urn:ietf:params:rtp-hdrext:ssrc-audio-level a=extmap:2 http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time a=extmap:3 http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01 a=extmap:4 urn:ietf:params:rtp-hdrext:sdes:mid a=extmap:5 urn:ietf:params:rtp-hdrext:sdes:rtp-stream-id a=extmap:6 urn:ietf:params:rtp-hdrext:sdes:repaired-rtp-stream-id a=sendrecv a=msid:9Frw0xYyaszH0JHrspJgRLPEs6eczFMDpBtt 29553c2e-9439-4ea2-9e91-9812892338db a=rtcp-mux a=rtpmap:111 opus/48000/2 a=rtcp-fb:111 transport-cc a=fmtp:111 minptime=10;useinbandfec=1 a=rtpmap:103 ISAC/16000 a=rtpmap:104 ISAC/32000 a=rtpmap:9 G722/8000 a=rtpmap:0 PCMU/8000 a=rtpmap:8 PCMA/8000 a=rtpmap:106 CN/32000 a=rtpmap:105 CN/16000 a=rtpmap:13 CN/8000 a=rtpmap:110 telephone-event/48000 a=rtpmap:112 telephone-event/32000 a=rtpmap:113 telephone-event/16000 a=rtpmap:126 telephone-event/8000 a=ssrc:1118557228 cname:E4YVGHfEu6fh9lLe a=ssrc:1118557228 msid:9Frw0xYyaszH0JHrspJgRLPEs6eczFMDpBtt 29553c2e-9439-4ea2-9e91-9812892338db a=ssrc:1118557228 mslabel:9Frw0xYyaszH0JHrspJgRLPEs6eczFMDpBtt a=ssrc:1118557228 label:29553c2e-9439-4ea2-9e91-9812892338db m=video 9 UDP/TLS/RTP/SAVPF 96 97 98 99 100 101 102 121 127 120 125 107 108 109 35 36 124 119 123 118 114 115 116 c=IN IP4 0.0.0.0 a=rtcp:9 IN IP4 0.0.0.0 a=ice-ufrag:zh8p a=ice-pwd:95ldnP+KE1KcFetSAHJaqj5Z a=ice-options:trickle a=fingerprint:sha-256 82:DB:EA:2B:C4:BA:BD:30:55:E5:A4:56:21:EE:B5:92:93:D6:CB:BA:DB:E5:DD:05:B9:F3:97:BC:AE:48:2C:98 a=setup:actpass a=mid:1 a=extmap:14 urn:ietf:params:rtp-hdrext:toffset a=extmap:2 http://www.webrtc.org/experiments/rtp-hdrext/abs-send-time a=extmap:13 urn:3gpp:video-orientation a=extmap:3 http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01 a=extmap:12 http://www.webrtc.org/experiments/rtp-hdrext/playout-delay a=extmap:11 http://www.webrtc.org/experiments/rtp-hdrext/video-content-type a=extmap:7 http://www.webrtc.org/experiments/rtp-hdrext/video-timing a=extmap:8 http://www.webrtc.org/experiments/rtp-hdrext/color-space a=extmap:4 urn:ietf:params:rtp-hdrext:sdes:mid a=extmap:5 urn:ietf:params:rtp-hdrext:sdes:rtp-stream-id a=extmap:6 urn:ietf:params:rtp-hdrext:sdes:repaired-rtp-stream-id a=sendrecv a=msid:9Frw0xYyaszH0JHrspJgRLPEs6eczFMDpBtt e0e4fa48-2931-4fef-b55d-8222dd2b9fe1 a=rtcp-mux a=rtcp-rsize a=rtpmap:96 VP8/90000 a=rtcp-fb:96 goog-remb a=rtcp-fb:96 transport-cc a=rtcp-fb:96 ccm fir a=rtcp-fb:96 nack a=rtcp-fb:96 nack pli a=rtpmap:97 rtx/90000 a=fmtp:97 apt=96 a=rtpmap:98 VP9/90000 a=rtcp-fb:98 goog-remb a=rtcp-fb:98 transport-cc a=rtcp-fb:98 ccm fir a=rtcp-fb:98 nack a=rtcp-fb:98 nack pli a=fmtp:98 profile-id=0 a=rtpmap:99 rtx/90000 a=fmtp:99 apt=98 a=rtpmap:100 VP9/90000 a=rtcp-fb:100 goog-remb a=rtcp-fb:100 transport-cc a=rtcp-fb:100 ccm fir a=rtcp-fb:100 nack a=rtcp-fb:100 nack pli a=fmtp:100 profile-id=2 a=rtpmap:101 rtx/90000 a=fmtp:101 apt=100 a=rtpmap:102 H264/90000 a=rtcp-fb:102 goog-remb a=rtcp-fb:102 transport-cc a=rtcp-fb:102 ccm fir a=rtcp-fb:102 nack a=rtcp-fb:102 nack pli a=fmtp:102 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42001f a=rtpmap:121 rtx/90000 a=fmtp:121 apt=102 a=rtpmap:127 H264/90000 a=rtcp-fb:127 goog-remb a=rtcp-fb:127 transport-cc a=rtcp-fb:127 ccm fir a=rtcp-fb:127 nack a=rtcp-fb:127 nack pli a=fmtp:127 level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=42001f a=rtpmap:120 rtx/90000 a=fmtp:120 apt=127 a=rtpmap:125 H264/90000 a=rtcp-fb:125 goog-remb a=rtcp-fb:125 transport-cc a=rtcp-fb:125 ccm fir a=rtcp-fb:125 nack a=rtcp-fb:125 nack pli a=fmtp:125 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=42e01f a=rtpmap:107 rtx/90000 a=fmtp:107 apt=125 a=rtpmap:108 H264/90000 a=rtcp-fb:108 goog-remb a=rtcp-fb:108 transport-cc a=rtcp-fb:108 ccm fir a=rtcp-fb:108 nack a=rtcp-fb:108 nack pli a=fmtp:108 level-asymmetry-allowed=1;packetization-mode=0;profile-level-id=42e01f a=rtpmap:109 rtx/90000 a=fmtp:109 apt=108 a=rtpmap:35 AV1X/90000 a=rtcp-fb:35 goog-remb a=rtcp-fb:35 transport-cc a=rtcp-fb:35 ccm fir a=rtcp-fb:35 nack a=rtcp-fb:35 nack pli a=rtpmap:36 rtx/90000 a=fmtp:36 apt=35 a=rtpmap:124 H264/90000 a=rtcp-fb:124 goog-remb a=rtcp-fb:124 transport-cc a=rtcp-fb:124 ccm fir a=rtcp-fb:124 nack a=rtcp-fb:124 nack pli a=fmtp:124 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=4d001f a=rtpmap:119 rtx/90000 a=fmtp:119 apt=124 a=rtpmap:123 H264/90000 a=rtcp-fb:123 goog-remb a=rtcp-fb:123 transport-cc a=rtcp-fb:123 ccm fir a=rtcp-fb:123 nack a=rtcp-fb:123 nack pli a=fmtp:123 level-asymmetry-allowed=1;packetization-mode=1;profile-level-id=64001f a=rtpmap:118 rtx/90000 a=fmtp:118 apt=123 a=rtpmap:114 red/90000 a=rtpmap:115 rtx/90000 a=fmtp:115 apt=114 a=rtpmap:116 ulpfec/90000 a=ssrc-group:FID 936231622 2179045404 a=ssrc:936231622 cname:E4YVGHfEu6fh9lLe a=ssrc:936231622 msid:9Frw0xYyaszH0JHrspJgRLPEs6eczFMDpBtt e0e4fa48-2931-4fef-b55d-8222dd2b9fe1 a=ssrc:936231622 mslabel:9Frw0xYyaszH0JHrspJgRLPEs6eczFMDpBtt a=ssrc:936231622 label:e0e4fa48-2931-4fef-b55d-8222dd2b9fe1 a=ssrc:2179045404 cname:E4YVGHfEu6fh9lLe a=ssrc:2179045404 msid:9Frw0xYyaszH0JHrspJgRLPEs6eczFMDpBtt e0e4fa48-2931-4fef-b55d-8222dd2b9fe1 a=ssrc:2179045404 mslabel:9Frw0xYyaszH0JHrspJgRLPEs6eczFMDpBtt a=ssrc:2179045404 label:e0e4fa48-2931-4fef-b55d-8222dd2b9fe1",
			Type: "offer",
		},
		// Candidates: []model.IceCandidate {
		// 	 {
		// 		Candidate: "candidate:3132842025 1 udp 2113937151 5435236b-caf4-48d1-90c5-6dbcbedfa3c7.local 52886 typ host generation 0 ufrag zh8p network-cost 999",
		// 		SdpMLineIndex: 1,
		// 		SdpMid: 0,
		// 	},
		// },
	}
	model.CreateWebRTCConnection(mockPeerCon)

}
