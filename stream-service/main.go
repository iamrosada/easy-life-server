package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pion/webrtc/v3"
)

// Room representa uma sala de conferência
type Room struct {
	ID        string
	Users     []string
	Messages  []Message
	CreatedAt time.Time
}

// Message representa uma mensagem enviada na sala de chat
type Message struct {
	User    string    `json:"user"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}

// WebRTCConnection representa uma conexão WebRTC em uma sala
type WebRTCConnection struct {
	PeerConnection *webrtc.PeerConnection
}

// Rooms mantém o estado de todas as salas
var Rooms = make(map[string]*Room)
var WebRTCRooms = make(map[string][]*WebRTCConnection)

// CriarSala cria uma nova sala
func CriarSala(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	salaID := params["id"]

	if _, existe := Rooms[salaID]; existe {
		http.Error(w, "Sala já existe", http.StatusBadRequest)
		return
	}

	sala := &Room{
		ID:        salaID,
		Users:     []string{},
		Messages:  []Message{},
		CreatedAt: time.Now(),
	}

	Rooms[salaID] = sala

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Sala criada com sucesso"))
}

// ObterSala retorna informações sobre uma sala
func ObterSala(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	salaID := params["id"]

	if sala, existe := Rooms[salaID]; existe {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(sala)
	} else {
		http.Error(w, "Sala não encontrada", http.StatusNotFound)
	}
}

// AdicionarUsuario adiciona um usuário a uma sala
func AdicionarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	salaID := params["id"]
	user := params["user"]

	if sala, existe := Rooms[salaID]; existe {
		sala.Users = append(sala.Users, user)
		w.Write([]byte("Usuário adicionado à sala"))
	} else {
		http.Error(w, "Sala não encontrada", http.StatusNotFound)
	}
}

// EnviarMensagem envia uma mensagem para a sala de chat
func EnviarMensagem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	salaID := params["id"]
	user := params["user"]

	if sala, existe := Rooms[salaID]; existe {
		var mensagem Message
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&mensagem); err != nil {
			http.Error(w, "Erro ao decodificar a mensagem", http.StatusBadRequest)
			return
		}

		mensagem.User = user
		mensagem.Created = time.Now()

		sala.Messages = append(sala.Messages, mensagem)

		w.Write([]byte("Mensagem enviada com sucesso"))
	} else {
		http.Error(w, "Sala não encontrada", http.StatusNotFound)
	}
}

// HandleOffer lida com a oferta SDP do cliente
func HandleOffer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roomID := params["id"]

	var offer webrtc.SessionDescription
	if err := json.NewDecoder(r.Body).Decode(&offer); err != nil {
		http.Error(w, "Erro ao decodificar a oferta", http.StatusBadRequest)
		return
	}

	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{},
	}

	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		http.Error(w, "Erro ao criar a conexão do Peer", http.StatusInternalServerError)
		return
	}

	connection := &WebRTCConnection{
		PeerConnection: peerConnection,
	}

	connection.PeerConnection.OnICECandidate(func(candidate *webrtc.ICECandidate) {
		if candidate != nil {

		}
	})

	connection.PeerConnection.OnDataChannel(func(d *webrtc.DataChannel) {
	})

	WebRTCRooms[roomID] = append(WebRTCRooms[roomID], connection)

	// Criar a resposta SDP
	answer, err := connection.PeerConnection.CreateAnswer(nil)
	if err != nil {
		http.Error(w, "Erro ao criar a resposta SDP", http.StatusInternalServerError)
		return
	}

	// Configurar a descrição da sessão no PeerConnection
	if err := connection.PeerConnection.SetLocalDescription(answer); err != nil {
		http.Error(w, "Erro ao definir a descrição local", http.StatusInternalServerError)
		return
	}

	// Enviar a resposta SDP ao cliente
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(answer)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/sala/{id}", CriarSala).Methods("POST")
	r.HandleFunc("/sala/{id}", ObterSala).Methods("GET")
	r.HandleFunc("/sala/{id}/usuario/{user}", AdicionarUsuario).Methods("POST")
	r.HandleFunc("/sala/{id}/mensagem/{user}", EnviarMensagem).Methods("POST")

	r.HandleFunc("/offer/{id}", HandleOffer).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
