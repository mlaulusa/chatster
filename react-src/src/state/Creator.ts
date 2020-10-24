import Action, { TYPE } from './Action'
import Message from '../Message'

export default class Creator {

  static addMessage (user: string, message: string, room: string = 'global') {
    return new Action(TYPE.ADD_MESSAGE, new Message(user, message, room))
  }

  static receiveMessage (message: Message) {
    return new Action(TYPE.RECEIVE_MESSAGE, message)
  }

  static setRooms (rooms: Array<Room>) {
    return new Action(TYPE.SET_ROOMS, rooms)
  }

}