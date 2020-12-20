import QuotesList from "components/QuotesList";
import { ChangeEvent, FormEvent, useEffect, useState } from "react";
import { QuoteResponse } from "types";
import Image from "next/image";

const IndexPage: React.FC = () => {
  const [search, setSearch] = useState<string>("");
  const [data, setData] = useState<QuoteResponse>();
  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    setSearch(event.target.value);
  };

  const handleSubmit = async (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const options = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ query: search }),
    };
    getData(options);
  };

  const getData = async (options?: any) => {
    const res =
      options == undefined
        ? await fetch(`http://localhost:9000/search`)
        : await fetch(`http://localhost:9000/search`, options);
    const resjson = await res.json();
    console.log(resjson);
    setData(resjson);
  };

  useEffect(() => {
    getData();
  }, []);

  return (
    <>
      <h2 className="text-center text-4xl text-indigo-900 font-display font-semibold lg:text-left xl:text-5xl xl:text-bold">
        Search for Quotes
      </h2>
      <div className="pt-2 relative mx-auto text-gray-600">
        <form onSubmit={handleSubmit}>
          <input
            className="border-2 border-gray-300 bg-white h-14 w-full px-5 pr-16 rounded-full text-md text-indigo-800 focus:outline-none"
            type="text"
            name="search"
            value={search}
            onChange={handleChange}
            placeholder="Search"
          />
          <button type="submit" className="absolute right-0 top-0 mt-5 mr-4">
            <Image
              src="/search-512.png"
              alt="Search icon"
              width={30}
              height={30}
            />
          </button>
        </form>
      </div>
      {data && data.quotes && (
        <>
          <div className="text-right text-md text-indigo-400 font-italic">
            {data.totalHits} results in {data.milliseconds} ms
          </div>
          <div>
            <QuotesList quotes={data.quotes}></QuotesList>
          </div>
        </>
      )}
    </>
  );
};

export default IndexPage;
