import { Tab } from "@headlessui/react";
import { PlusSmIcon } from "@heroicons/react/solid";
import { formatRFC3339, parse } from "date-fns";
import { useEffect, VFC } from "react";
import { Link } from "react-router-dom";
import { PartialDeep } from "type-fest";
import {
  GetClassesBySchoolQuery,
  useGetClassesBySchoolLazyQuery,
} from "../generated/graphql";

type Props = {
  date: string;
  onClickClass: (clazz: ClassType) => void;
  onClickNewClass?: (schedule: PartialDeep<SchedulesType[number]>) => void;
};

type StudiosType = GetClassesBySchoolQuery["school"]["studios"];
type StudioType = StudiosType[number];
type RoomsType = StudioType["rooms"];
type RoomType = RoomsType[number];
type SchedulesType = RoomType["schedules"];
type ClassType = NonNullable<SchedulesType[number]["class"]>;
type InstructorType = ClassType["instructor"];

export const Timetable: VFC<Props> = ({ date, onClickClass }) => {
  const [getClassesBySchool, { data, loading, error }] =
    useGetClassesBySchoolLazyQuery();

  useEffect(() => {
    getClassesBySchool({
      variables: {
        id: "1",
        date: formatRFC3339(parse(date, "yyyyMMdd", new Date())),
      },
    });
  }, [date, getClassesBySchool]);

  const renderTabGroupForStudios = (studios: StudiosType) => (
    <Tab.Group>
      <Tab.List>
        {studios.map((studio) => (
          <Tab key={studio.id}>{studio.name}</Tab>
        ))}
      </Tab.List>
      <Tab.Panels>
        {studios.map((studio) => (
          <Tab.Panel key={studio.id}>
            {renderTabGroupForRooms(studio.rooms)}
          </Tab.Panel>
        ))}
      </Tab.Panels>
    </Tab.Group>
  );

  const renderTabGroupForRooms = (rooms: RoomsType) => (
    <Tab.Group>
      <Tab.List>
        {rooms.map((room) => (
          <Tab key={room.id}>{room.name}</Tab>
        ))}
      </Tab.List>
      <Tab.Panels>
        {rooms.map((room) => (
          <Tab.Panel key={room.id}>{renderSchedules(room.schedules)}</Tab.Panel>
        ))}
      </Tab.Panels>
    </Tab.Group>
  );

  const renderSchedules = (schedules: SchedulesType) => {
    return (
      <div className="grid grid-cols-8">
        {[0, 1, 2, 3, 4, 5, 6, 7].map((dayOfWeek) => {
          return <div key={dayOfWeek}>{dayOfWeek}</div>;
        })}
        {renderTimeslots(schedules)}
      </div>
    );
  };

  const renderTimeslots = (schedules: SchedulesType) => {
    return [
      { startTime: "13:00", endTime: "14:00" },
      { startTime: "14:15", endTime: "15:25" },
      { startTime: "15:30", endTime: "16:40" },
      { startTime: "16:45", endTime: "17:55" },
      { startTime: "18:00", endTime: "19:10" },
      { startTime: "19:15", endTime: "20:25" },
      { startTime: "20:30", endTime: "21:40" },
      { startTime: "21:45", endTime: "22:55" },
    ].map((slot) => {
      return [0, 1, 2, 3, 4, 5, 6, 7].map((dayOfWeek) => {
        if (dayOfWeek === 0) {
          return (
            <div
              key={dayOfWeek}
              className="text-md flex flex-col px-4 py-2 font-bold text-gray-700"
            >
              <div>{slot.startTime}</div>
              <div>~ {slot.endTime}</div>
            </div>
          );
        }
        const schedule = schedules.find(
          (schedule) =>
            schedule.dayOfWeek === dayOfWeek &&
            schedule.startTime === slot.startTime
        );
        if (!schedule) return <div key={`${dayOfWeek}-none`}></div>;
        if (schedule.class) {
          return (
            <div key={`${dayOfWeek}-${schedule.id}`}>
              {renderClass(schedule.class)}
            </div>
          );
        } else {
          return (
            <div key={`${dayOfWeek}-${schedule.id}`}>
              {renderNewClassButton(schedule)}
            </div>
          );
        }
      });
    });
  };

  const renderClass = (clazz: ClassType) => {
    if (!clazz) return <></>;
    return (
      <div onClick={() => onClickClass(clazz)}>
        <div>{clazz.name}</div>
        <div>{clazz.level}</div>
        {renderInstructor(clazz.instructor)}
      </div>
    );
  };

  const renderInstructor = (instructor: InstructorType) => {
    return <div className="text-md text-gray-700">{instructor.name}</div>;
  };

  const renderNewClassButton = (
    schedule?: PartialDeep<SchedulesType[number]>
  ) => {
    if (!schedule) return <></>;
    return (
      <Link to={`/schedules/${schedule.id}/classes/new`}>
        <PlusSmIcon className="h-4 w-4" />
      </Link>
    );
  };

  if (!data) {
    return <></>;
  }
  const {
    school: { studios },
  } = data;

  return (
    <>
      <h1>クラス</h1>
      <div className="bg-white">{renderTabGroupForStudios(studios)}</div>
    </>
  );
};
