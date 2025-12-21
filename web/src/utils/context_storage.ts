import type { ContextStorageItem, ContextStorageItemPair } from "../types";

class ContextStorage {
    private storage: Map<string, ContextStorageItem<any>>;
    public constructor () {
        this.storage = new Map();
    }
    public get<T>(key: string): ContextStorageItem<T> | undefined {
        return this.storage.get(key);
    }
    public set<T>(items: ContextStorageItemPair<T>[]): void {
        items.forEach((item: ContextStorageItemPair<T>) => this.storage.set(item.key, item.contextItem))
    }
}

export default ContextStorage;