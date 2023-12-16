import React from "react";
import { Chart } from "./chart";

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
                <div className="h-[20rem] rounded-lg bg-gray-50 rounded-lg border border-gray-300  bg-gray-700  border-gray-600 ">
                    <Chart elements={file} />
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