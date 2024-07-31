import {Table} from "react-bootstrap";

export default function EncounterRow({ m }) {
    return (
        <Table striped bordered className="table-sm mb-3">
            <tbody>
            <tr>
                <th style={{width: "15%"}}>Name</th>
                <th>Weak</th>
                <th>Null</th>
                <th>Absorb</th>
                <th>Flags</th>
            </tr>
            <tr key={m.name}>
                <td rowSpan={7}>
                    <strong>{m.name}</strong>
                </td>
                <td>
                    {m.elementWeak ? m.elementWeak?.join(", ") : "-"}
                </td>
                <td>
                    {m.elementNull ? m.elementNull?.join(", ") : "-"}
                </td>
                <td>
                    {m.elementAbsorb ? m.elementAbsorb?.join(", ") : "-"}
                </td>
                <td>
                    {m.flags ? m.flags?.join(", ") : "-"}
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
                    m.metamorphItems ? m.metamorphItems.join(", ").replaceAll("\\pad", "") : "-"
                }</td>
                <td>{m.commonSteal != null ? m.commonSteal.replace("\\pad", "") : "None"} {m.rareSteal != null && ", " + m.rareSteal.replace("\\pad", "")}</td>
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

            <tr>
                <th colSpan={1}>Sketch</th>
                <th colSpan={1}>Rage</th>
                <th colSpan={2}>Control</th>
            </tr>

            <tr>
                <td>{m.sketch1 && m.sketch1} {m.sketch2 && ", " + m.sketch2}</td>
                <td>{m.rage1 && m.rage1} {m.rage2 && ", " + m.rage2}</td>
                <td colSpan={2}>{m.control1 && m.control1} {m.control2 && ", " +  m.control2} {m.control3 && ", " +  m.control3} {m.control4 && ", " + m.control4}</td>
            </tr>
            </tbody>
        </Table>
    )
}
