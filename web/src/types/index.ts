export type WebSocketMsg = {
    type: string
    content: string
}

export type ContextStorageItem<T> = {
    value: T
    setValue: React.Dispatch<React.SetStateAction<T>>,
}

export type ContextStorageItemPair<T> = {
    key: string,
    contextItem: ContextStorageItem<T>,
}