# simple messenger initial

this is a simple messenger.

#### 1. AddUser
```
this function get two inputs:
1. username : length of username should be greater than 3 and must be unique and must contain letter and digit.
2. IsBot : it is a bool that determine typr of a user. 
if inputes be correct the function will return a ID for user other wise return the error.
```
#### 2. AddChat
```
this function get four input : 
1. chatname : this is name of the chat
2. isGroup : thios is a bool that detemine type of the chat. channel or group
3. creator : this is a int that show the owner of chat ID.
4. admins : this is a rray of int that implement the ID's of admins.
if all inputs be correct function will return the ID of chat other wise return the error.
```

#### 3. SendMessage
```
this function get three inputs:
userId : the id of sender of message.
chatId : the id of chat that we send the message.
text : text of message.
Note : if type of chat be channel only admins have prmision to send message and if the user not be admin function will return error.
if every thing be correct the function will return ID of message.
```

#### 4. SendLike
``` get two inputs:
1. userId: the user that wants like a message.
2. messageId : ID of the message that user want to like.
if the message doese not exist function will return message not found error.
Note: everu user can like once a message not more.
if a user like a message for second or more time function will return error. 
```
#### 5. GetNumberOfLikes
```
this function get messageId as input and will return number of likes.
```

#### 6. SetChatAdmin
```
this function get two inputs:
1. chatId : Id of the chat that we want add a admin to it.
2. userId : Id of new admin.
Note : if the user already exist in admins function will return error.
```

#### 7. GetLastMessage
```
this function get id of a chat and will return last message of the chat by massageId and text.
```
#### 8. GetLastUserMessage
```
this function get id of a user and will return last message of that user by massageId and text.
```