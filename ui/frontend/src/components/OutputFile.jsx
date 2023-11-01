import React from "react";

function Output(props) {
    return (
        <div>
            <label
                htmlFor="output"
                className="pt-5 block mb-2 text-lg font-medium text-gray-900 dark:text-white"
            >
                Output
            </label>
            <div className="py-2.5">
                <textarea
                    id="output"
                    rows="8"
                    className="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                />
            </div>
        </div>
    );
}
export default Output;
