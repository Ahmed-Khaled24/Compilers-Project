import React from "react";
import Cytoscape from "cytoscape";
import COSEBilkent from "cytoscape-cose-bilkent";
import CytoscapeComponent from "react-cytoscapejs";

Cytoscape.use(COSEBilkent);

export const Chart = (props) => {
  const layout = {
    name: "cose-bilkent",
    padding:10,
    edgeElasticity: 0.1,
    height:'700px'
  };

  const cytoscapeStylesheet = [
    {
      selector: "node",
      style: {
        "background-color": "green",
        width: "label",
        height: "label",
        padding: "6px",
        shape: "circle"
      }
    },
    
    {
      selector: 'node[label = "op"],node[label = "On"]',
      style: {
        shape: "rectangle"
      }
    },
    {
      selector: "node[label]",
      style: {
        label: "data(label)",
        "font-size": "12",
        color: "white",
        "text-halign": "center",
        "text-valign": "center"
      }
    },
    {
      selector: "edge",
      style: {
        width: 1
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

      }}    stylesheet={cytoscapeStylesheet}
  />
    
  );
};
