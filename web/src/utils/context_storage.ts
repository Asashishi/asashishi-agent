export type ContextStorageItemPair<T> = {
    key: string,
    contextItem: ContextStorageItem<T>,
};

export class ContextStorageItem<T> {
    public value: T;
    private setState: React.Dispatch<React.SetStateAction<T>>;

    public constructor(
        value: T,
        setState: React.Dispatch<React.SetStateAction<T>>
    ) {
        this.value = value;
        this.setState = setState;
    }

    public setValue(nValue: T): void {
        this.value = nValue;
        this.setState(nValue);
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