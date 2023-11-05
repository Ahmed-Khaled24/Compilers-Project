import "./App.css";
import React, { useState } from "react";
import Dropdown from "./components/Dropdown";
import InputFile from "./components/InputFile";
import OutputFile from "./components/OutputFile";
import { Scan } from "../wailsjs/go/Scanner/ScannerStruct";

function Input() {
    const [selection, setSelection] = useState("");
    const [scannerResult, setScannerResult] = useState("No results yet.");
    const [inputFile, setInputFile] = useState("");

    const options = [
        { label: "Scanner", value: "scanner" },
        { label: "Parser", value: "parser" },
    ];

    const handleSelect = (option) => {
        setSelection(option);
    };

    const handleFileChange = (e) => {
        const file = e.target.files[0];
        const reader = new FileReader();
        reader.readAsText(file);
        reader.onload = () => {
            setInputFile(reader.result);
        };
        reader.onerror = () => {
            console.log("file error", reader.error);
        };
    };

    const clear = () => {
        const fileInput = document.getElementById("file-input");
        if (fileInput) {
            fileInput.value = "";
        }
        setInputFile("");
       // setScannerResult("No results yet.");
    };
    const handleSubmit = async (e) => {
        e.preventDefault();
        let analysis = await Scan(inputFile);
        analysis = analysis.map((token) => {
            let { TokenBaseType, ...rest } = token;
            return rest;
        });
        if (analysis.length === 0) {
            setScannerResult("No results yet.");
        } else {
            const formattedOutput = JSON.stringify(analysis, null, 4);
            setScannerResult(formattedOutput);
        }
    };
    function SaveOutput(scannerResult) {
        const blob = new Blob([scannerResult], { type: 'text/plain' });
        const url = URL.createObjectURL(blob);
        const link = document.createElement('a');
        link.download = 'Result.txt';
        link.href = url;
        link.click();
    }


    return (
        <form onSubmit={handleSubmit}>
            <div className="flex justify-around py-8 w-full border bg-gray-800 border-gray-600 min-h-screen">
                <div className="w-4/6 px-4  rounded-t-lg bg-gray-800 pb-18">
                    <div className="flex items-center justify-center space-x-6 px-auto">
                        <div className="w-1/4">
                            <Dropdown
                                options={options}
                                value={selection}
                                onChange={handleSelect}
                            />
                        </div>

                        <div className=" h-full w-2/5">
                            <input
                                id="file-input"
                                type="file"
                                className="relative m-0 block w-full min-w-0 flex-auto rounded border border-solid border-neutral-300 bg-clip-padding px-3 py-[0.32rem] text-base font-normal text-neutral-700 transition duration-300 ease-in-out file:-mx-3 file:-my-[0.32rem] file:overflow-hidden file:rounded-none file:border-0 file:border-solid file:border-inherit file:bg-neutral-100 file:px-3 file:py-[0.32rem] file:text-neutral-700 file:transition file:duration-150 file:ease-in-out file:[border-inline-end-width:1px] file:[margin-inline-end:0.75rem] file:bg-blue-200 focus:border-primary focus:text-neutral-700 focus:shadow-te-primary focus:outline-none  border-neutral-600  text-neutral-200  file:bg-blue-700  file:text-white  focus:border-primary text-white file:bg-sky-800"
                                onChange={handleFileChange}
                            />
                        </div>

                        <div className=" h-full w-1/6 pt-2">
                            <button
                                type="submit"
                                className="text-white font-medium rounded-lg text-sm px-5 py-2 mr-1 mb-2  bg-green-600  hover:bg-green-700 focus:outline-none  focus:ring-green-800 w-full"
                            >
                                Analyze
                            </button>
                        </div>
                    </div>
                    <InputFile file={inputFile} change={setInputFile} />
                    <OutputFile file={scannerResult} />
                    <div className="   py-3 flex flex-row-reverse">
                        <button
                            class="w-1/5 text-white bg-red-600 hover:bg-red-700  font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2"
                            onClick={clear}>
                            Clear
                        </button>
                        <button
                            class="w-1/5 text-white bg-sky-800 hover:bg-blue-900  font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2"
                            onClick={() => SaveOutput(scannerResult)}>
                            Save
                        </button>
                    </div>
                </div>
            </div>
        </form>
    );
}

function App() {
    return (
        <div>
            <Input />
        </div>
    );
}

export default App;
