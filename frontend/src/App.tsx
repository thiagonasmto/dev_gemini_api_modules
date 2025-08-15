import { useState, useRef } from "react";
import { sendMessage, type GeminiRequestMessage } from "./services/geminiService";
import TextareaAutosize from "react-textarea-autosize";
import ReactMarkdown from "react-markdown";
import "./App.css";

function App() {
  const [message, setMessage] = useState("");
  const [messages, setMessages] = useState<GeminiRequestMessage[]>([]);
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const TypingIndicator = () => {
    return (
      <div className="typing-indicator">
        <span></span>
        <span></span>
        <span></span>
      </div>
    );
  };

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (!e.target.files || e.target.files.length === 0) return;
    setSelectedFile(e.target.files[0]);
    console.log("Arquivo carregado")
  };
  
  const handleSend = async () => {
    if (!message && !selectedFile) return;

    // Adiciona mensagem do usuário
    const userMessage: GeminiRequestMessage = {
      role: "user",
      content: message,
      file: selectedFile || undefined,
    };

    setMessages((prev) => [...prev, userMessage]);

    // Adiciona placeholder de "digitando..."
    const loadingMessage: GeminiRequestMessage = {
      role: "model",
      content: "__TYPING__",
    };
    setMessages((prev) => [...prev, loadingMessage]);

    try {
      const responseText = await sendMessage(userMessage);

     setMessages((prev) => {
        const newMessages = [...prev];
        newMessages.pop();
        return [...newMessages, { role: "model", content: responseText }];
      });
    } catch {
      setMessages((prev) => {
        const newMessages = [...prev];
        newMessages.pop();
        return [...newMessages, { role: "model", content: "Erro ao processar requisição" }];
      });
    }

    setMessage("");
    setSelectedFile(null); 
    if (fileInputRef.current) {
      fileInputRef.current.value = "";
    }
  };

  return (
    <div className="container-chat">
      <div className="chat-messages">
        {messages.map((m, i) => (
          <div key={i} className={`message-bubble ${m.role === "user" ? "user" : "model"}`} >
            {m.content === "__TYPING__" ? <TypingIndicator /> : <ReactMarkdown>{m.content}</ReactMarkdown>}
          </div>
        ))}
      </div>

      <div className="prompt-input">
        <TextareaAutosize
          minRows={1}
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          placeholder="Digite sua mensagem"
          onKeyDown={(e) => {
            if (e.key === "Enter" && !e.shiftKey) {
              e.preventDefault();
              handleSend();
            }
          }}
        />

        <div className="inputs">
          <input
            type="file"
            ref={fileInputRef}
            style={{ display: "none" }}
            onChange={handleFileChange}
          />

          <button className="file-input" onClick={() => fileInputRef.current?.click()}>📎</button>
          {/* <button className="text-input" onClick={handleSend}>Enviar</button> */}
        </div>
      </div>
    </div>
  );
}

export default App;