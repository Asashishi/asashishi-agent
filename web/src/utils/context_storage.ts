export type ContextStorageItemPair<T> = {
    key: string,
    contextItem: ContextStorageItem<T>,
}


export class ContextStorageItem<T> {
    public value: T;
    public setValue: (item: T) => void;
    public constructor(value: T, setValue: React.Dispatch<React.SetStateAction<T>>) {
        this.value = value;
        this.setValue = (next: T): void => {
            setValue(next);
            this.value = next;
        } 
    }
};

class ContextStorage {
    private storage: Map<string, ContextStorageItem<any>>;
    public constructor () {
        this.storage = new Map();
    }
    public get<T>(key: string): ContextStorageItem<T> | undefined {
        return this.storage.get(key);
    }
    public set<T>(items: ContextStorageItemPair<T>[]): void {
        items.forEach((item: ContextStorageItemPair<T>) => this.storage.set(item.key, item.contextItem));
    }
}

export default ContextStorage;