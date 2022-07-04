import axios from "axios";
import moment from "moment";
import getConfig from "next/config";
import { CustomHoliday } from "types";

export const getCustomHolidays = async () => {
  const { publicRuntimeConfig } = getConfig();
  const response = await axios.get(
    `${publicRuntimeConfig.apiURL}/api/v1/custom-holiday`
  );
  return response.data.map((datum: any) => ({
    ...datum,
    date: moment(datum.date).format("YYYY-MM-DD"),
  }));
};

export const addCustomHoliday = async ({ date, category }: CustomHoliday) => {
  const { publicRuntimeConfig } = getConfig();
  const response = await axios.post(
    `${publicRuntimeConfig.apiURL}/api/v1/custom-holiday`,
    {
      date,
      category,
    }
  );

  return {
    date,
    category,
  };
};

export const deleteCustomHoliday = async (date: string) => {
  const { publicRuntimeConfig } = getConfig();
  const response = await axios.delete(
    `${publicRuntimeConfig.apiURL}/api/v1/custom-holiday`,
    {
      params: {
        date,
      },
    }
  );

  return date;
};

export const updateCustomHoliday = async ({
  date,
  category,
}: CustomHoliday) => {
  const { publicRuntimeConfig } = getConfig();
  const response = await axios.put(
    `${publicRuntimeConfig.apiURL}/api/v1/custom-holiday`,
    {
      date,
      category,
    }
  );

  return date;
};
