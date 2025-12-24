import React, { createContext, useState, type JSX, type ReactNode } from "react";
import type { DisplayMsg, Tab } from "../types/websocket_type";
import wsInstance from "../utils/websocket";

export type ContextItems<T> = Record<string, T | React.Dispatch<React.SetStateAction<T>>>

export const AsashishiAgentContext = createContext<ContextItems<any>>({});

const AppContext = ({ children }: { children: ReactNode }): JSX.Element => {
    const [tab, setTab] = useState<Tab>("chat");
    const [ioHistories, setIOHistories] = useState<DisplayMsg[]>([]);
    const context: ContextItems<any> = {
        tab,
        setTab,
        ioHistories,
        setIOHistories,
    };
    wsInstance.injectDependencies([
        {
            key: "tab",
            value: tab,
            setValue: setTab,
        },
        {
            key: "ioHistories",
            value: ioHistories,
            setValue: setIOHistories,
        },
    ]);
    return (
        <AsashishiAgentContext.Provider value={context}>
            {children}
        </AsashishiAgentContext.Provider>
    )
}

export default AppContext;