import { StatusRoute } from "../consts/http"
import type { Tab } from "../types/websocket_type"

const syncServerStaus: (action: React.Dispatch<React.SetStateAction<Tab>>) => Promise<void> = async (action: React.Dispatch<React.SetStateAction<Tab>>): Promise<void> => {
    const tab: Tab = await fetch(StatusRoute).then((result) => result.text()).then((text) => text.trim()) as Tab
    return action(tab);
};

export default syncServerStaus;