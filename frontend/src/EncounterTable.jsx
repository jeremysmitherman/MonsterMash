import {Table} from "react-bootstrap";
import {useEffect, useRef, useState} from "react";
import {GetCurrentEncounter, GetClientStatus } from "../wailsjs/go/ff6library/Library.js";
import EncounterRow from "./EncounterRow.jsx";

export default function EncounterTable() {
    const [enemies, setEnemies] = useState([])
    const [clientState, setClientState] = useState({status: "Unknown", status_message: ""})
    const encounterTimer = useRef(0)

    useEffect(() => {
        clearInterval(encounterTimer.current);
        encounterTimer.current = setInterval(async () => {
            let newEncounter = await GetCurrentEncounter()
            let newState = await GetClientStatus()
            setEnemies(newEncounter)

            if(newState.status !== clientState.status) {
                setClientState(newState);
            }

            return () => { console.log("Clearing interval"); clearInterval(encounterTimer.current); }
        }, 1800)
    }, [enemies, clientState, encounterTimer]);

    let element = <h3 style={{color: "white" }}>Waiting for encounter...</h3>

    if (enemies && enemies != []) {
        element = <>
            {enemies.map((e) => (<EncounterRow m={e}></EncounterRow> ))}
        </>
    }

    return (
        <>
            <div className="mb-5">
                { element }
            </div>
            <p className={(clientState.status === "Connected") ? "mt-3 text-white" : "mt-3 text-danger"}>Status: {clientState.status}</p>
            <p>
                <small>{clientState.status_message}</small>
            </p>
        </>)
}