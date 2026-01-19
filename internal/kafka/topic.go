package kafka

const defaultTopic = "default_message"

// topic is shared by producer/consumer within this process.
// It is set during initialization and then treated as read-only.
var topic = defaultTopic

func SetTopic(topicInput string) {
	if topicInput != "" {
		topic = topicInput
	}
}
