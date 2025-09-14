package main 

func main() {
	pub1 := NewPublisher()
	pub2 := NewPublisher()

	topic1 := NewTopic("Topic1")
	topic2 := NewTopic("Topic2")

	sub1 := NewOrderSubscriber("sub1")
	sub2 := NewOrderSubscriber("sub2")
	sub3 := NewOrderSubscriber("sub3")
	sub4 := NewOrderSubscriber("sub4")
	sub5 := NewOrderSubscriber("sub5")

	topic1.AddSubscriber(sub1)
	topic1.AddSubscriber(sub2)
	topic1.AddSubscriber(sub3)
	topic1.AddSubscriber(sub4)
	topic1.AddSubscriber(sub5)

	topic2.AddSubscriber(sub1)
	topic2.AddSubscriber(sub3)
	topic2.AddSubscriber(sub4)


	pub1.RegisterTopic(topic1)
	pub1.RegisterTopic(topic2)

	pub2.RegisterTopic(topic2)

	m1 := NewMessage("m1", "This is m1")
	m2 := NewMessage("m2", "This is m2")
	m3 := NewMessage("m3", "This is m3")
	m4 := NewMessage("m4", "This is m4")


	pub1.Publish(topic1, m1)
	pub1.Publish(topic2, m2)


	pub2.Publish(topic2, m3)
	pub2.Publish(topic1, m3)

	topic1.RemoveSubscriber(sub5)
	
	pub1.Publish(topic1, m4)

}