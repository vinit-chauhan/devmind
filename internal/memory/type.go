package memory

import (
	"fmt"
	"os"
	"strings"

	"github.com/vinit-chauhan/devmind/internal/agent/types"
	"github.com/vinit-chauhan/devmind/internal/logger"
	"github.com/vinit-chauhan/devmind/internal/memory/chat"
)

var Brain *Memory = nil

type Memory struct {
	Chats []chat.Chat
}

func New() *Memory {
	return &Memory{
		Chats: []chat.Chat{},
	}
}

func (m *Memory) AddChat(role, content, createdAt string) {
	m.Chats = append(m.Chats, chat.Chat{
		Role:      role,
		Content:   content,
		CreatedAt: createdAt,
	})
}

func (m *Memory) LoadMemories(chats string) {
	counter := 0
	for line := range strings.SplitSeq(chats, SPLITTER) {
		chat := chat.Chat{}

		if err := chat.Parse(line); err != nil {
			logger.Error("Error parsing chat: " + err.Error() + " line: " + line)
			continue
		}

		counter++
		m.Chats = append(m.Chats, chat)
	}
	logger.Debug(fmt.Sprintf("Loaded %d chats from memory\n", counter))
}

func (m *Memory) AddChatToMemory(conversation []chat.Chat) {
	for _, chat := range conversation {
		m.Chats = append(m.Chats, chat)
	}
	m.SaveToMemory(conversation)
}

func (m *Memory) SaveToMemory(conversation []chat.Chat) error {
	// Open the file for writing
	file, err := os.OpenFile(MEMORY_FILE, os.O_APPEND, 0644)
	if err != nil {
		logger.Error("Error opening memory file: " + err.Error())
	}
	defer file.Close()

	for _, chat := range conversation {
		if _, err := file.WriteString(chat.String() + "\n"); err != nil {
			logger.Error("Error writing to memory file: " + err.Error())
			return err
		}
	}

	file.WriteString(SPLITTER)

	return nil
}

func (m *Memory) GetLast(n int) []chat.Chat {
	if n > len(m.Chats) {
		n = len(m.Chats)
	}
	return m.Chats[len(m.Chats)-n:]
}

func (m *Memory) GetChatMessageHistory() []types.Message {
	chatHistory := m.GetLast(HISTORY_LENGTH)

	messages := []types.Message{}
	for _, chat := range chatHistory {
		messages = append(messages, types.Message{
			Role:    chat.Role,
			Content: chat.Content,
		})
	}
	return messages
}

func (m *Memory) GetMemoryPrompt() string {
	chatHistory := m.GetLast(HISTORY_LENGTH)

	if len(chatHistory) > 0 {
		prompt := strings.Builder{}

		prompt.WriteString("Chat History:\n")
		for _, chat := range chatHistory {
			prompt.WriteString(chat.String())
			prompt.WriteString("\n")
		}
		return prompt.String()
	}
	return ""
}

func init() {
	// Check if the memory file exists
	if _, err := os.Stat(MEMORY_FILE); os.IsNotExist(err) {
		// Create the file if it doesn't exist
		file, err := os.Create(MEMORY_FILE)
		if err != nil {
			logger.Error("Error creating memory file:" + err.Error())
			return
		}
		defer file.Close()
	}

	contents, err := os.ReadFile(MEMORY_FILE)
	if err != nil {
		logger.Error("Error reading memory file:" + err.Error())
		return
	}

	Brain = New()

	if len(contents) != 0 {
		logger.Info("Loading memory from file")
		Brain.LoadMemories(string(contents))
		logger.Info("Memory loaded successfully")
	}

	logger.Info("Memory initialized successfully")
}
