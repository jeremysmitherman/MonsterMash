import {Table} from "react-bootstrap";
import {useEffect, useRef, useState} from "react";
import {GetCurrentEncounter, GetClientStatus } from "../wailsjs/go/reference/Library.js";

export default function EncounterTable() {
    const [encounter, setEncounter] = useState({encounter_id: -1, monsters: []})
    const [clientState, setClientState] = useState({status: "Unknown", status_message: ""})
    const encounterTimer = useRef(0)

    useEffect(() => {
        clearInterval(encounterTimer.current);
        encounterTimer.current = setInterval(async () => {
            console.log("Checking for new encounter...")
            let newEncounter = await GetCurrentEncounter()
            let newState = await GetClientStatus()

            if(encounter.encounter_id !== newEncounter.encounter_id) {
                setEncounter(newEncounter);
            }

            if(newState.status !== clientState.status) {
                console.log("Setting new state: " + newState.status + " from: " + clientState.status)
                setClientState(newState);
            }

            return () => { console.log("Clearing interval"); clearInterval(encounterTimer.current); }
        }, 1800)
    }, [encounter, clientState, encounterTimer]);

    return (
        <div className="enemyTable mb-5">
            <p style={{color: "white"}}>{encounter.encounter_id} - { encounter.monsters.length }</p>
            {clientState.status === "Connected" && encounter.monsters?.map((m) => {
                return (
                    <Table striped bordered className="table-sm mb-0">
                        <tbody>
                        <tr>
                            <th>Name</th>
                            <th>Weak</th>
                            <th>Null</th>
                            <th>Absorb</th>
                            <th>Flags</th>
                        </tr>
                        <tr key={m.name}>
                            <td rowSpan={5}>
                                <strong>{m.name}</strong>
                            </td>
                            <td>
                                {m.elementWeak?.join(", ")}
                            </td>
                            <td>
                                {m.elementNull?.join(", ")}
                            </td>
                            <td>
                                {m.elementAbsorb?.join(", ")}
                            </td>
                            <td>
                                {m.flags?.join(", ")}
                            </td>
                        </tr>

                        <tr>
                            <th colSpan={1}>Initial Status</th>
                            <th colSpan={1}>Immune</th>
                            <th colSpan={1}>Metamorph ({m.morphRate})</th>
                            <th colSpan={1}>Steal (Common, Rare)</th>
                        </tr>

                        <tr>
                            <td colSpan={1}> &nbsp; {m.statusSet?.join(", ")}</td>
                            <td colSpan={1}> &nbsp; {m.statusImmunity?.join(", ")}</td>
                            <td colSpan={1}> &nbsp; {
                                m.metamorphItems.join(", ").replaceAll("\\pad", "")
                            }</td>
                            <td>{m.commonSteal != null && m.commonSteal.replace("\\pad", "")} {m.rareSteal != null && ", " + m.rareSteal.replace("\\pad", "")}</td>
                        </tr>

                        <tr>
                            <th>XP</th>
                            <th>Gil</th>
                            <th>Common Drop</th>
                            <th>Rare Drop</th>
                        </tr>

                        <tr>
                            <td>{m.experience}</td>
                            <td>{m.gp}</td>
                            <td>{m.commonDrop !== "\\padEmpty" && m.commonDrop.replace("\\pad", "")}</td>
                            <td>{m.rareDrop !== "\\padEmpty" && m.rareDrop.replace("\\pad", "")}</td>
                        </tr>
                        </tbody>
                    </Table>
                )
            })}
                <h3 className="mt-5 text-white">Status: {clientState.status}</h3>
                <p>
                    <small>{clientState.status_message}</small>
                </p>
            </div>)
}