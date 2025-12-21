import type { ContextStorageItem, ContextStorageItemPair } from "../types";

class ContextStorage {
    private storage: Map<string, ContextStorageItem<any>>;
    constructor () {
        this.storage = new Map();
    }
    public get<T>(key: string): ContextStorageItem<T> | undefined {
        return this.storage.get(key);
    }
    public set<T>(items: ContextStorageItemPair<T>[]): void {
        for (const item of items) {
            this.storage.set(item.key, item.contextItem);
        }
    }
}

export default ContextStorage;