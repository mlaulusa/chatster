export function hasKeys<S extends Object> (object: Object, keys: Array<string>): object is S {
  return Object.keys(object).some(keys.includes.bind(keys))
}