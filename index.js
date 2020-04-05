var proc =  require('child_process');
var lo = require('lodash');
var proj = process.argv[2];
var result = proc.execSync(`./stat ${proj}`).toString();
var commits = result.split("---------------------------");

const fileInfos ={};
commits.forEach(commit => {
    var body = commit.split("*******")
    var changesBody = body[3]
    
    const commitInfo = {
        author: (body[1] || "").replace('\n',''),
        commitHash: body[0],
        commitmessage: (body[2] || "")
    }
    
    if(changesBody){
        var filesStat = changesBody.split("\n");
        filesStat.forEach(fileStat=>{
            if(fileStat.replace(" ","")!==""){
                const fileChanges = fileStat.split("\t");
                const fileName = fileChanges[2];

                const fileAdds = parseInt(fileChanges[0]);
                const fileRemotions = parseInt(fileChanges[1]);

                const fileInfo  = fileInfos[fileName] || {
                    commits:[],
                    totalAdds: 0,
                    totalRemotions: 0,
                    totalChanges: 0
                }
                fileInfos[fileName] = {
                    commits: [...fileInfo.commits, commitInfo],
                    totalAdds: fileInfo.totalAdds + (isNaN(fileAdds) ? 0 : fileAdds),
                    totalRemotions: fileInfo.totalRemotions + (isNaN(fileRemotions) ? 0 : fileRemotions),
                    totalChanges: fileInfo.totalChanges + fileRemotions + fileAdds
                }
            }
        })
    }
});


lo(fileInfos)
    .map((fileInfo,fileName)=>({...fileInfo,fileName}))
    .orderBy(['totalChanges'],['desc'])
    .take(50)
    .value()
    .forEach(({ fileName, totalChanges, commits}) => console.log(`fileName: ${fileName}\ has ${totalChanges} changes in ${commits.length} commits by ${lo.uniq(commits.map(c=>c.author))}\n`));