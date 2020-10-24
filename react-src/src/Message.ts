import { DateTime } from 'luxon'
import { hasKeys } from './Utils'

interface messageJson {
  user: string
  payload: string
  room: string
  time: string
}

export default class Message implements messageJson {

  static fromJSON (json: Object): Message {

    if (hasKeys<messageJson>(json, ['user', 'payload', 'room', 'time'])) {
      return new Message(json.user, json.payload, json.room, json.time)
    } else {
      throw TypeError('object does not have all keys (user, payload, room, time)')
    }
  }

  constructor (private _user: string, private _payload: string, private _room: string, private _isoDateTime: string = DateTime.local().toISO()) {
  }

  toJSON = (): messageJson => ({
    user: this._user,
    payload: this._payload,
    room: this._room,
    time: this._isoDateTime,
  })

  get user () {
    return this._user
  }

  get payload () {
    return this._payload
  }

  get room () {
    return this._room
  }

  get time (): string {
    return this._isoDateTime
  }

  get dateTime (): DateTime {
    return DateTime.fromISO(this._isoDateTime)
  }

}

