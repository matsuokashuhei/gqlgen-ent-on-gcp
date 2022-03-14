import { useEffect, VFC } from "react";
import { Layout } from "../../components";
import { GetSchoolQuery, useGetSchoolLazyQuery } from "../../generated/graphql";

export const ClassesPage: VFC = () => {
  const [getSchool, { data, loading, error }] = useGetSchoolLazyQuery();

  useEffect(() => {
    getSchool({ variables: { id: "1" } });
  }, [getSchool]);

  const renderStudios = (studios: GetSchoolQuery["school"]["studios"]) => {
    return studios.map((studio) => renderStudio(studio));
  };

  const renderStudio = (
    studio: GetSchoolQuery["school"]["studios"][number]
  ) => {
    return (
      <div key={studio.id}>
        {studio.name}
        {renderRooms(studio.rooms)}
      </div>
    );
  };

  const renderRooms = (
    rooms: GetSchoolQuery["school"]["studios"][0]["rooms"]
  ) => {
    return rooms.map((room) => renderRoom(room));
  };

  const renderRoom = (
    room: GetSchoolQuery["school"]["studios"][0]["rooms"][0]
  ) => {
    return (
      <div key={room.id}>
        {room.name}
        {renderSchedules(room.schedules)}
      </div>
    );
  };

  const renderSchedules = (
    schedules: GetSchoolQuery["school"]["studios"][number]["rooms"][number]["schedules"]
  ) => {
    return (
      <div className="grid grid-cols-8">
        {[0, 1, 2, 3, 4, 5, 6, 7].map((dayOfWeek) => {
          return <div key={dayOfWeek}>{dayOfWeek}</div>;
        })}
        {renderTimeslots(schedules)}
      </div>
    );
  };

  const renderTimeslots = (
    schedules: GetSchoolQuery["school"]["studios"][number]["rooms"][number]["schedules"]
  ) => {
    return [
      "13:00",
      "14:15",
      "15:30",
      "16:45",
      "18:00",
      "19:15",
      "20:30",
      "21:45",
    ].map((startTime) => {
      return [0, 1, 2, 3, 4, 5, 6, 7].map((dayOfWeek, index) => {
        if (dayOfWeek === 0) {
          return <div key={dayOfWeek}>{startTime}</div>;
        }
        const schedule = schedules.find(
          (schedule) =>
            schedule.dayOfWeek === dayOfWeek && schedule.startTime === startTime
        );
        if (schedule) {
          return (
            <div key={schedule.id}>
              {schedule.startTime}
              {schedule.endTime}
            </div>
          );
        } else {
          return <div>None</div>;
        }
      });
    });
  };

  if (!data) {
    return <></>;
  }
  const {
    school: { studios },
  } = data;

  return (
    <Layout>
      <h1>クラス</h1>
      {renderStudios(studios)}
    </Layout>
  );
};
