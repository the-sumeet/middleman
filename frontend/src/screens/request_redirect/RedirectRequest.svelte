<div class="w-full p-6 mx-auto rounded-md shadow-md bg-gray-800">
    <h2 class="text-lg font-semibold capitalize text-white">
        If
    </h2>

    <div class="mb-4 pl-6">
        <div class="flex flex-col gap-6 mt-4 sm:grid-cols-2">
            <div class="flex space-x-2">
                <!-- ENTITY -->
                <select
                    class="w-min bg-white border text-sm rounded-lg  block p-3 bg-gray-800 border-gray-600 placeholder-gray-400 text-white focus:ring-blue-500 focus:border-blue-500"
                >
                    <option selected>URL</option>
                    <option selected>Query</option>
                    <option value="US">Header</option>
                    <option value="CA">Body</option>
                </select>

                <!-- AS -->
                <label
                    class="p-2 text-gray-200"
                    for="username">As</label
                >

                <!-- TYPE -->
                <select
                    class="w-min border  text-sm rounded-lg  block p-3 bg-gray-800 border-gray-600 placeholder-gray-400 text-white focus:ring-blue-500 focus:border-blue-500"
                >
                    <option selected>String</option>
                    <option selected>Number</option>
                    <option value="US">List</option>
                </select>

                <!-- IS -->
                <label
                    class="p-2 text-gray-200"
                    for="username">Is</label
                >

                <!-- OP -->
                <select
                    class="w-min border  text-sm rounded-lg  block p-3 bg-gray-800 border-gray-600 placeholder-gray-400 text-white focus:ring-blue-500 focus:border-blue-500"
                >
                    <option selected>Equal</option>
                    <option selected>Contain</option>
                    <option selected>Less than</option>
                    <option value="US">Less than or Equal</option>
                    <option selected>Greater than</option>
                    <option value="US">Greater than</option>
                </select>
            </div>

            <div class="">
                
                <input
                    type="text"
                    class="block w-full px-4 py-2 text-gray-700  border border-gray-200 rounded-md bg-gray-800 text-gray-300 border-gray-600 focus:border-blue-400 focus:ring-blue-300 focus:ring-opacity-40 focus:border-blue-300 focus:outline-none focus:ring"
                />
            </div>
        </div>
    </div>

    <h2 class="text-lg font-semibold  capitalize text-white">
        Then
    </h2>

    <div class="mb-4 pl-6">
        <div class="flex flex-col gap-6 mt-4 sm:grid-cols-2">
            <div class="">
                <!-- Redirect To -->
                <label
                    class=" text-gray-200"
                    for="username">Redirect To</label
                >
                <input
                    type="text"
                    placeholder="https://example.com"
                    class="block w-full px-4 py-2   border  rounded-md bg-gray-800 text-gray-300 border-gray-600  focus:ring-blue-300 focus:ring-opacity-40 focus:border-blue-300 focus:outline-none focus:ring"
                />
            </div>
        </div>
    </div>

    <div class="flex justify-end mt-6">
        <button
            class="px-8 py-2.5 leading-5 text-white transition-colors duration-300 transform bg-gray-700 rounded-md hover:bg-gray-600 focus:outline-none focus:bg-gray-600"
            >Save</button
        >
    </div>
</div>
