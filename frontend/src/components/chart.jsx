import React from "react";
import Cytoscape from "cytoscape";
import COSEBilkent from "cytoscape-cose-bilkent";
import CytoscapeComponent from "react-cytoscapejs";

Cytoscape.use(COSEBilkent);

export const Chart = (props) => {
    const layout = {
        name: "cose-bilkent",
        componentSpacing: 60,
        fit: true,
        padding: 10,
        edgeElasticity: 2,
        height: '800px',
        width: "800px",
        avoidOverlap: true


    };
    const cytoscapeStylesheet = [
        {
            selector: "node",
            style: {
                "background-color": "#1F2937",
                width: "label",
                minWidth: "60px",
                height: "label",
                padding: "16px",
                shape: "circle"
            }
        },

        {
            selector: 'node[label ^= "assign"],node[label^= "read"],node[label^= "Repeat"],node[label^= "write"] ,node[label^= "if"]',
            style: {
                shape: "rectangle",
                padding: "14px",

            }
        },
        {
            selector: "node[label]",
            style: {
                label: "data(label)",
                "font-size": "12",
                color: "#DDDFE1",
                "text-halign": "center",
                "text-valign": "center"
            }
        },
        {
            selector: "edge",
            style: {
                width: 1,
                
            }
        },
        {
            selector: 'edge[label="Child"]',
            style: {
                width: 3,
                'line-color': "#1F2937",

            }
        },
        {
            selector: 'edge[label="Next"]',
            style: {
                width: 3,
                'line-color': "#DC2626",
            }
        },
    ];

    return (
        <CytoscapeComponent

            elements={props.elements}
            layout={layout}
            style={{
                top: 0,
                buttom: 0,
                width: "100%",
                height: "100%",

            }}
            stylesheet={cytoscapeStylesheet}
        />

    );
};
