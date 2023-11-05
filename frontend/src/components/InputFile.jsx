import React from "react";

function Input({ file, change }) {
    return (
        <div>
            <label
                htmlFor="input"
                className="pt-10 block mb-2 text-lg font-medium text-gray-900  text-white uppercase"
            >
                Input
            </label>
            <div className="py-2.5">
                <textarea
                    value={file}
                    id="input"
                    rows={10}
                    className="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500  bg-gray-700  border-gray-600  placeholder-gray-400  text-white  focus:ring-blue-500  focus:border-blue-500"
                    onChange={(e) => change(e.target.value)}
                ></textarea>
            </div>
        </div>
    );
}
export default Input;
