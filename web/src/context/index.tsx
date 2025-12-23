import React, { createContext, useState, type JSX, type ReactNode } from "react";
import type { DisplayMsg } from "../types/websocket_type";

export type Tab = "AI" | "Shell";
export type ContextItems<T> = Record<string, T | React.Dispatch<React.SetStateAction<T>>>

export const AsashishiAgentContext = createContext<ContextItems<any>>({});

const AppContext = ({ children }: { children: ReactNode }): JSX.Element => {
    const [tab, setTab] = useState<Tab>("AI");
    const [ioHistories, setIOHistories] = useState<DisplayMsg[]>([]);
    const context: ContextItems<any> = {
        tab,
        setTab,
        ioHistories,
        setIOHistories,
    }
    return (
        <AsashishiAgentContext.Provider value={context}>
            {children}
        </AsashishiAgentContext.Provider>
    )
}

export default AppContext;