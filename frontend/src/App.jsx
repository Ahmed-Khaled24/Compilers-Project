import "./App.css";
import React, { useState } from "react";
import Dropdown from "./components/Dropdown";
import InputFile from "./components/InputFile";
import OutputFile from "./components/OutputFile";
import { Scan } from "../wailsjs/go/Scanner/ScannerStruct";
import { Parse } from "../wailsjs/go/Parser/Parser";
function Input() {
    const [selection, setSelection] = useState({
        label: "Scanner",
        value: "scanner",
    });
    const [scannerResult, setScannerResult] = useState("No results yet.");
    const [inputFile, setInputFile] = useState("");
    const [parserResult, setParserResult] = useState("");

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
        setParserResult("");
        // setScannerResult("No results yet.");
    };
    const y=[]
    const handleSubmit = async (e) => {
        e.preventDefault();
        let filtered = inputFile.replaceAll(/[\n\r\t]/g, " ");
        let analysis = await Scan(filtered);
        const Result = await Parse(filtered);
        analysis = analysis.map((token) => {
            let { TokenBaseType, ...rest } = token;
            return rest;
        });
        if (analysis.length === 0) {
            setScannerResult("No results yet.");
        } else {
            console.log(Result);
            setParserResult(renderNode(Result).concat(y))
            const formattedOutput = JSON.stringify(analysis, null, 4);
            setScannerResult(formattedOutput);
        }
    };

    function SaveOutput(scannerResult) {
        const blob = new Blob([scannerResult], { type: "text/plain" });
        const url = URL.createObjectURL(blob);
        const link = document.createElement("a");
        link.download = "Result.txt";
        link.href = url;
        link.click();
    }
    let nodeId = 0;

    function renderNode(node) {
        const currentId = `${nodeId++}`;
        const x = [
            {
                data: {
                    id: currentId,
                    label: node.NodeType + '  ( ' + node.NodeValue + ' )',
                },
            }
        ];

        if (node.Children) {
            node.Children.map((child) => {
                const childNodes = renderNode(child);
                    x.push(...(childNodes));
                y.push({
                    data: {
                        source: currentId,
                        target: childNodes[0].data.id,
                        label: 'Child',

                    },
                });
            });
        }
        if (node.Next) {
            const nextNodes = renderNode(node.Next);
                x.push(...(nextNodes));
            y.push({
                data: {
                    source: currentId,
                    target: nextNodes[0].data.id,
                    label: 'Next',
                },
            });
        }

        let ans=x.filter((value,index) => x.indexOf(value)==index)
        y.sort((a,b)=>{
            let x = parseInt(a.data.source);
            let y = parseInt(b.data.source);
            if(x>y){return 1;}
            if(x<y){return -1;}
            return 0;
          });
        return ans;
    }

    return (
        <form onSubmit={handleSubmit}>
            <div className="flex justify-around py-8 w-full border bg-gray-800 border-gray-600 min-h-screen">
                <div className="w-5/6 px-4  rounded-t-lg bg-gray-800 pb-18">
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
                    <OutputFile
                        file={
                            selection.value === "scanner"
                                ? scannerResult
                                : parserResult
                        }
                        type={selection}
                    />
                    <div className="   py-3 flex flex-row-reverse">
                        <button
                            class="w-1/5 text-white bg-red-600 hover:bg-red-700  font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2"
                            onClick={clear}
                        >
                            Clear
                        </button>
                        <button
                            class="w-1/5 text-white bg-sky-800 hover:bg-blue-900  font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2"
                            onClick={() => SaveOutput(scannerResult)}
                        >
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
