import { useState, useEffect } from 'react'

export default function Home() {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);
  useEffect(() => {
    getData()
  }, []);

  const getData = async () => {
    try {
      const res = await fetch(process.env.NEXT_PUBLIC_BACKEND_URL);
      const data = await res.json();
      setData(data);
    }
    catch (error) {
      setError(error);
    }
  }

  const handleCheck = async (e) => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/updateCheck/${e.currentTarget.value}`, {
        method: 'PATCH',
      });
      getData();
    }
    catch (error) {
      setError(error);
    }
  }

  const handleDelete = async (e) => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/${e.currentTarget.value}`, {
        method: 'DELETE',
      });
      getData();
    }
    catch (error) {
      setError(error);
    }
  }

  return (
    <div className='container mx-10 my-7'>
      {error && <div>Failed to load {error.toString()}</div>}
      {
        !data ? <div>Loading...</div>
          : (
            (data?.data ?? []).length === 0 && <p>data kosong</p>
          )
      }

      <Input onSuccess={getData} />
      {data?.data && data?.data?.map((item, index) => (
        <div key={index}>
          <input type="checkbox" defaultChecked={item.done} onChange={handleCheck} value={item.ID} />
          <span> {item.ID}. Task: {item.task}</span>
          <button value={item.ID} onClick={handleDelete} className="bg-red-500 rounded-lg">Delete</button>
        </div>
      ))}
    </div>
  )
}

function Input({ onSuccess }) {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const body = {
      task: formData.get("data")
    }

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/send`, {
        method: 'POST',
        body: JSON.stringify(body)
      });
      const data = await res.json();
      setData(data.message);
      onSuccess();
    }
    catch (error) {
      setError(error);
    }
  }

  return (
    <div>
      {error && <p>error: {error.toString()}</p>}
      {data && <p>success: {data}</p>}
      <form onSubmit={handleSubmit}>
        <input name="data" type="text" />
        <button className='bg-green-500 rounded-lg'>Submit</button>
      </form>
    </div>
  )
}
