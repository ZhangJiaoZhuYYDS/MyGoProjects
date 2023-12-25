// @Author zhangjiaozhu 2023/12/24 17:09:00
package main

import (
	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
	"log"
	"net/http"
)

// WebSocket升级器
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 处理WebSocket连接
func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	// 创建WebRTC配置
	config := webrtc.Configuration{}

	// 创建WebRTC API
	api := webrtc.NewAPI(webrtc.WithMediaEngine(&webrtc.MediaEngine{}))

	// 创建WebRTC PeerConnection
	peerConnection, err := api.NewPeerConnection(config)
	if err != nil {
		log.Println("Failed to create PeerConnection:", err)
		return
	}

	// 处理ICE候选事件
	peerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate != nil {
			log.Println("ICE candidate:", candidate)
			err := conn.WriteJSON(candidate.ToJSON())
			if err != nil {
				log.Println("Failed to send ICE candidate:", err)
			}
		}
	})

	// 处理SDP事件
	peerConnection.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		log.Println("ICE connection state:", connectionState)
	})

	// 从客户端接收和处理消息
	go func() {
		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				log.Println("WebSocket read error:", err)
				break
			}

			// 处理接收到的SDP消息
			// 这里假设客户端发送的是offer SDP消息
			offer := webrtc.SessionDescription{}
			//err, _ = offer.Unmarshal(msg)
			//if err != nil {
			//	log.Println("Failed to unmarshal offer:", err)
			//	break
			//}

			// 设置远端描述
			err = peerConnection.SetRemoteDescription(offer)
			if err != nil {
				log.Println("Failed to set remote description:", err)
				break
			}

			// 创建应答
			answer, err := peerConnection.CreateAnswer(nil)
			if err != nil {
				log.Println("Failed to create answer:", err)
				break
			}

			// 设置本地描述
			err = peerConnection.SetLocalDescription(answer)
			if err != nil {
				log.Println("Failed to set local description:", err)
				break
			}

			// 发送应答给客户端
			err = conn.WriteJSON(answer)
			if err != nil {
				log.Println("Failed to send answer:", err)
				break
			}
		}
	}()

	// 交换ICE候选
	go func() {
		for {
			var candidate webrtc.ICECandidateInit
			err := conn.ReadJSON(&candidate)
			if err != nil {
				log.Println("Failed to read ICE candidate:", err)
				break
			}

			// 添加ICE候选到PeerConnection
			err = peerConnection.AddICECandidate(candidate)
			if err != nil {
				log.Println("Failed to add ICE candidate:", err)
				break
			}
		}
	}()

	// 等待通话结束
	select {}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Println("Server listening on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
}
