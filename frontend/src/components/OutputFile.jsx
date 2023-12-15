import React from "react";
import { Chart } from "./chart";
const basicElements = [
    { data: { id: "read", label: "read" } },
    { data: { id: "if", label: "if" } },
    { data: { id: "op", label: "op" } },
    { data: { id: "assign", label: "assign" } },
    { data: { id: "const", label: "const" } },
    { data: { id: "id", label: "id" } },
    { data: { id: "const1", label: "const" } },
    { data: { id: "repeat", label: "repeat" } },
    { data: { id: "op1", label: "op" } },
    { data: { source: "assign", target: "repeat" } },
    { data: { source: "assign", target: "const1" } },   
    { data: { source: "repeat", target: "op1" } },
    { data: { source: "read", target: "if" } },
    { data: { source: "if", target: "op" } },
    { data: { source: "if", target: "assign" } },
    { data: { source: "op", target: "id" } },
    { data: { source: "op", target: "const" } },
   
];


function Output({ file, type }) {
    return (
        <div>
            <label
                htmlFor="output"
                className="pt-5 block mb-2 text-lg font-medium text-gray-900 text-white uppercase"
            >
                output
            </label>
            {type.value === "parser" ? (
                // <div className="py-2 h[600px] bg-gray-700 rounded-lg border border-gray-800">
                //     <div className="tree" >
                //         {treeRendering(treeData)}
                //     </div>
                // </div>

                <div className="h-[20rem] rounded-lg ">
                    <Chart elements={basicElements} />
                </div>
            ) : (
                <div className="py-2.5">
                    <textarea
                        id="output"
                        rows={15}
                        className="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 bg-gray-700 border-gray-600 placeholder-gray-400 text-white focus:ring-blue-500 focus:border-blue-500"
                        value={file}
                        disabled={true}
                    />
                </div>
            )}
        </div>
    );
}

export default Output;