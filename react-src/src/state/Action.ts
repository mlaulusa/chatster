export default class Action {
  payload: Object

  constructor (private _type: string, payload: Object) {
    this.payload = payload
  }

  get type () {
    return this._type
  }
}

export const TYPE = {
  ADD_MESSAGE: 'ADD_MESSAGE',
  RECEIVE_MESSAGE: 'RECEIVE_MESSAGE',
}
