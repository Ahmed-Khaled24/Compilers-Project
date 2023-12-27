import React from "react";
import CytoscapeComponent from "react-cytoscapejs";
import cytoscape from 'cytoscape';
import dagre from 'cytoscape-dagre';

cytoscape.use( dagre ); // register extension


export const Chart = (props) => {
    const layout = {
        name: "dagre",
        fit: true,
        padding: 10,
        edgeElasticity: 1,
        height: '800px',
        width: "800px",
        direction:true,
        avoidOverlap: true,
        NodeSpacing:3,
        randomize: false,           
  
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
                shape: "circle",


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
            selector: 'node[id = "0"]',
            style: {
                shape: "rectangle",
                "background-color": "#16A34A",
                padding: "14px",

            }
        },
        {
            selector: "node[label]",
            style: {
                label: "data(label)",
                "font-size": "14",
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
